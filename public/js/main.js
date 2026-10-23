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

  // links from home open a preselected tab via ?tab=…; keep the address bar
  // at plain /pengalaman so tabs don't pile up as separate visited pages
  if (location.pathname === "/pengalaman" && location.search) {
    history.replaceState(null, "", "/pengalaman");
  }

  observe(document);
  // content swapped in by htmx (projects "show more", pengalaman tabs)
  document.addEventListener("htmx:afterSettle", () => observe(document));
})();
