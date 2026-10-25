// Motion + nav behaviour. Everything here is progressive enhancement:
// without JS the page renders fully visible and links still work.
(() => {
  const reduce = matchMedia("(prefers-reduced-motion: reduce)").matches;
  const finePointer = matchMedia("(hover: hover) and (pointer: fine)").matches;

  // ---- scroll reveal ----
  const io = "IntersectionObserver" in window && !reduce
    ? new IntersectionObserver((entries) => {
        for (const e of entries) {
          if (e.isIntersecting) {
            e.target.classList.add("in");
            io.unobserve(e.target);
          }
        }
      }, { rootMargin: "0px 0px -8% 0px", threshold: 0.12 })
    : null;

  function observe(root) {
    root.querySelectorAll(".reveal:not(.in), .reveal-line:not(.in)").forEach((el) => {
      if (io) io.observe(el);
      else el.classList.add("in");
    });
  }

  // ---- nav: scrolled state + active section ----
  const nav = document.querySelector(".nav");
  const onScroll = () => nav && nav.classList.toggle("is-scrolled", scrollY > 24);
  addEventListener("scroll", onScroll, { passive: true });
  onScroll();

  const links = [...document.querySelectorAll("[data-nav]")];
  const sections = links
    .map((a) => document.getElementById(a.dataset.nav))
    .filter(Boolean);
  if (sections.length && "IntersectionObserver" in window) {
    const spy = new IntersectionObserver((entries) => {
      for (const e of entries) {
        if (!e.isIntersecting) continue;
        for (const a of links) {
          a.setAttribute("aria-current", String(a.dataset.nav === e.target.id));
        }
      }
    }, { rootMargin: "-45% 0px -50% 0px" });
    sections.forEach((s) => spy.observe(s));
  }

  // ---- mobile menu ----
  const toggle = document.getElementById("menu-toggle");
  const setMenu = (open) => {
    document.documentElement.classList.toggle("menu-open", open);
    toggle.setAttribute("aria-expanded", String(open));
    toggle.setAttribute("aria-label", open ? "Close menu" : "Open menu");
    document.body.style.overflow = open ? "hidden" : "";
  };
  if (toggle) {
    toggle.addEventListener("click", () =>
      setMenu(toggle.getAttribute("aria-expanded") !== "true"));
    document.querySelectorAll("#menu-panel a").forEach((a) =>
      a.addEventListener("click", () => setMenu(false)));
    addEventListener("keydown", (e) => {
      if (e.key === "Escape" && toggle.getAttribute("aria-expanded") === "true") {
        setMenu(false);
        toggle.focus();
      }
    });
    matchMedia("(min-width: 768px)").addEventListener("change", (m) => m.matches && setMenu(false));
  }

  // ---- pointer: hero parallax + project card glow (desktop only) ----
  if (finePointer && !reduce) {
    const layers = document.querySelectorAll("[data-depth]");
    let raf = 0;
    addEventListener("pointermove", (e) => {
      if (raf) return;
      raf = requestAnimationFrame(() => {
        raf = 0;
        const x = e.clientX / innerWidth - 0.5;
        const y = e.clientY / innerHeight - 0.5;
        layers.forEach((l) => {
          const d = +l.dataset.depth;
          l.style.transform = `translate3d(${x * d}px, ${y * d}px, 0)`;
        });
      });
    }, { passive: true });

    document.addEventListener("pointermove", (e) => {
      const card = e.target.closest && e.target.closest(".proj");
      if (!card) return;
      const r = card.getBoundingClientRect();
      card.style.setProperty("--mx", `${e.clientX - r.left}px`);
      card.style.setProperty("--my", `${e.clientY - r.top}px`);
    }, { passive: true });
  }

  // ---- projects: show all / show less without moving the viewport ----
  const list = document.getElementById("project-list");
  const more = document.getElementById("project-toggle");
  if (list && more) {
    more.addEventListener("click", () => {
      const open = !list.classList.contains("is-open");
      const before = more.getBoundingClientRect().top;
      list.classList.toggle("is-open", open);
      more.setAttribute("aria-expanded", String(open));
      more.querySelector("[data-more]").hidden = open;
      more.querySelector("[data-less]").hidden = !open;
      // collapsing pulls the button up; scroll by the same amount so it stays
      // under the cursor instead of the page jumping to another section
      if (!open) {
        const delta = more.getBoundingClientRect().top - before;
        if (delta) scrollBy({ top: delta, behavior: "instant" });
      }
    });
  }

  // ---- tech stack: auto-scroll that can be swiped (touch) or dragged (LMB) ----
  // The track holds the items twice, so wrapping by half its width is seamless.
  document.querySelectorAll(".marquee").forEach((row, i) => {
    let pos = 0;
    let held = false; // finger down or mouse dragging
    let hovered = false;
    let resumeAt = 0;
    const speed = 0.035 - i * 0.004; // px per ms; rows drift at different paces
    const half = () => row.scrollWidth / 2;
    const wrap = (x) => {
      const h = half();
      return h > 0 ? ((x % h) + h) % h : x;
    };
    const release = () => { held = false; resumeAt = performance.now() + 1500; };

    // touch: native scrolling, we only pause the autoplay around it
    row.addEventListener("touchstart", () => { held = true; }, { passive: true });
    row.addEventListener("touchend", release, { passive: true });
    row.addEventListener("touchcancel", release, { passive: true });

    // mouse: hold the left button and drag left/right
    let startX = 0;
    let startPos = 0;
    row.addEventListener("pointerdown", (e) => {
      if (e.pointerType !== "mouse" || e.button !== 0) return;
      held = true;
      startX = e.clientX;
      startPos = row.scrollLeft;
      try { row.setPointerCapture(e.pointerId); } catch {} // keep dragging outside the row
      row.classList.add("is-dragging");
    });
    row.addEventListener("pointermove", (e) => {
      if (!held || e.pointerType !== "mouse") return;
      pos = wrap(startPos - (e.clientX - startX));
      row.scrollLeft = pos;
    });
    const endDrag = (e) => {
      if (e.pointerType !== "mouse" || !row.classList.contains("is-dragging")) return;
      row.classList.remove("is-dragging");
      release();
    };
    row.addEventListener("pointerup", endDrag);
    row.addEventListener("pointercancel", endDrag);
    row.addEventListener("mouseenter", () => { hovered = true; });
    row.addEventListener("mouseleave", () => { hovered = false; });

    if (reduce) return; // still draggable/swipeable, just no autoplay

    let last = performance.now();
    const tick = (now) => {
      const dt = now - last;
      last = now;
      if (held || hovered || now < resumeAt) {
        pos = row.scrollLeft; // follow the user's swipe / drag / momentum
      } else {
        pos = wrap(pos + speed * dt);
        row.scrollLeft = pos;
      }
      requestAnimationFrame(tick);
    };
    requestAnimationFrame(tick);
  });

  // ---- certificates: snap carousel with buttons, dots, mouse drag, lightbox ----
  const track = document.getElementById("cert-track");
  if (track) {
    const slides = [...track.children];
    const dots = [...document.querySelectorAll(".cert-dot")];
    const prev = document.getElementById("cert-prev");
    const next = document.getElementById("cert-next");
    const current = document.getElementById("cert-current");
    let active = -1;

    const setActive = (i) => {
      if (i === active) return;
      active = i;
      slides.forEach((s, j) => s.classList.toggle("is-active", j === i));
      dots.forEach((d, j) => d.setAttribute("aria-current", String(j === i)));
      current.textContent = String(i + 1).padStart(2, "0");
      prev.disabled = i === 0;
      next.disabled = i === slides.length - 1;
    };
    // the active slide is the one whose left edge sits closest to the snap line
    const nearest = () => {
      const edge = track.getBoundingClientRect().left + parseFloat(getComputedStyle(track).scrollPaddingLeft);
      let best = 0;
      slides.forEach((s, j) => {
        if (Math.abs(s.getBoundingClientRect().left - edge) < Math.abs(slides[best].getBoundingClientRect().left - edge)) best = j;
      });
      // at the far right end the last slide can't reach the snap line
      if (track.scrollLeft + track.clientWidth >= track.scrollWidth - 2) best = slides.length - 1;
      return best;
    };
    const go = (i) => {
      i = Math.max(0, Math.min(slides.length - 1, i));
      const delta = slides[i].getBoundingClientRect().left - slides[active].getBoundingClientRect().left;
      track.scrollBy({ left: delta, behavior });
      setActive(i);
    };

    let scrollRaf = 0;
    track.addEventListener("scroll", () => {
      if (scrollRaf) return;
      scrollRaf = requestAnimationFrame(() => { scrollRaf = 0; setActive(nearest()); });
    }, { passive: true });
    prev.addEventListener("click", () => go(active - 1));
    next.addEventListener("click", () => go(active + 1));
    dots.forEach((d) => d.addEventListener("click", () => go(+d.dataset.index)));
    track.addEventListener("keydown", (e) => {
      if (e.key === "ArrowRight") { e.preventDefault(); go(active + 1); }
      if (e.key === "ArrowLeft") { e.preventDefault(); go(active - 1); }
    });

    // mouse: drag with LMB;
    // a drag must not also count as a click that opens the lightbox
    let startX = 0, startLeft = 0, dragging = false, moved = false;
    track.addEventListener("pointerdown", (e) => {
      if (e.pointerType !== "mouse" || e.button !== 0) return;
      dragging = true; moved = false;
      startX = e.clientX; startLeft = track.scrollLeft;
    });
    addEventListener("pointermove", (e) => {
      if (!dragging) return;
      const dx = e.clientX - startX;
      if (!moved && Math.abs(dx) > 5) { moved = true; track.classList.add("is-dragging"); }
      if (moved) track.scrollLeft = startLeft - dx;
    });
    addEventListener("pointerup", () => {
      if (!dragging) return;
      dragging = false;
      // re-enabling snap makes the browser settle on the closest slide
      if (moved) track.classList.remove("is-dragging");
    });

    const box = document.getElementById("cert-lightbox");
    const boxImg = box.querySelector("img");
    track.addEventListener("click", (e) => {
      const btn = e.target.closest(".cert-media");
      if (!btn) return;
      if (moved) { moved = false; return; }
      boxImg.src = btn.dataset.full;
      boxImg.alt = "Certificate: " + btn.dataset.title;
      box.showModal();
    });
    // clicking the backdrop (outside the image) closes it
    box.addEventListener("click", (e) => { if (e.target === box) box.close(); });

    setActive(0);
  }

  // ---- in-page anchors: smooth scroll without leaving #section in the URL ----
  const behavior = reduce ? "auto" : "smooth";
  document.addEventListener("click", (e) => {
    const a = e.target.closest && e.target.closest("a[href*='#']");
    if (!a || e.button !== 0 || e.metaKey || e.ctrlKey || e.shiftKey || e.altKey) return;
    const url = new URL(a.getAttribute("href"), location.href);
    if (url.origin !== location.origin || url.pathname !== location.pathname) return;
    if (url.hash === "#main") return; // skip link keeps its native focus jump
    e.preventDefault();
    const target = url.hash.length > 1 && document.getElementById(url.hash.slice(1));
    if (target) target.scrollIntoView({ behavior });
    else scrollTo({ top: 0, behavior });
  });
  // arriving from another page via /#section: land on it, then drop the hash
  if (location.hash) {
    addEventListener("load", () => {
      const target = document.getElementById(location.hash.slice(1));
      if (target) target.scrollIntoView({ behavior: "auto" });
      history.replaceState(null, "", location.pathname + location.search);
    });
  }

  // links from home open a preselected tab via ?tab=…; keep the address bar
  // at plain /pengalaman so tabs don't pile up as separate visited pages
  if (location.pathname === "/pengalaman" && location.search) {
    history.replaceState(null, "", "/pengalaman");
  }

  observe(document);
  // content swapped in by htmx (projects "show more", pengalaman tabs)
  document.addEventListener("htmx:afterSettle", () => observe(document));
})();
