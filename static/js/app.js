/* ============================================================
   EZY GAMING — App JavaScript
   Alpine.js handles most interactivity. This file covers:
   - Scroll-triggered animations
   - Animated stat counters
   - Active nav link highlighting
============================================================ */

document.addEventListener('DOMContentLoaded', () => {
  initScrollAnimations();
  initStatCounters();
  highlightActiveNav();
  initPS5IframeLazyLoad();
});

/* ── Scroll-triggered section reveal ────────────────────── */
function initScrollAnimations() {
  const sections = document.querySelectorAll('section');
  sections.forEach(el => el.classList.add('fade-in-section'));

  const observer = new IntersectionObserver(
    (entries) => {
      entries.forEach(entry => {
        if (entry.isIntersecting) {
          entry.target.classList.add('is-visible');
          observer.unobserve(entry.target);
        }
      });
    },
    { threshold: 0.1, rootMargin: '0px 0px -60px 0px' }
  );

  sections.forEach(el => observer.observe(el));
}

/* ── Animated number counters ────────────────────────────── */
function initStatCounters() {
  const counters = document.querySelectorAll('[data-count]');
  if (!counters.length) return;

  const observer = new IntersectionObserver(
    (entries) => {
      entries.forEach(entry => {
        if (entry.isIntersecting) {
          animateCounter(entry.target);
          observer.unobserve(entry.target);
        }
      });
    },
    { threshold: 0.5 }
  );

  counters.forEach(el => observer.observe(el));
}

function animateCounter(el) {
  const target = parseInt(el.dataset.count, 10);
  const duration = 1800;
  const start = performance.now();

  function update(now) {
    const elapsed = now - start;
    const progress = Math.min(elapsed / duration, 1);
    // Ease out cubic
    const eased = 1 - Math.pow(1 - progress, 3);
    el.textContent = Math.floor(eased * target);
    if (progress < 1) requestAnimationFrame(update);
    else el.textContent = target;
  }

  requestAnimationFrame(update);
}

/* ── PS5 Sketchfab iframe — lazy load on scroll ──────────── */
function initPS5IframeLazyLoad() {
  const iframe = document.getElementById('ps5-controller-iframe');
  const placeholder = document.getElementById('ps5-iframe-placeholder');
  if (!iframe) return;

  const observer = new IntersectionObserver(
    (entries) => {
      entries.forEach(entry => {
        if (entry.isIntersecting) {
          // Inject src only when user scrolls into view
          if (!iframe.src && iframe.dataset.src) {
            iframe.src = iframe.dataset.src;

            // Fade out loading placeholder once iframe fires load
            iframe.addEventListener('load', () => {
              if (placeholder) {
                placeholder.style.transition = 'opacity 0.6s ease';
                placeholder.style.opacity = '0';
                setTimeout(() => placeholder.remove(), 600);
              }
            }, { once: true });
          }
          observer.unobserve(iframe);
        }
      });
    },
    { threshold: 0.2 }
  );

  observer.observe(iframe);
}

/* ── Active nav link ─────────────────────────────────────── */
function highlightActiveNav() {
  const path = window.location.pathname;
  document.querySelectorAll('.nav-link').forEach(link => {
    const href = link.getAttribute('href');
    const isActive = href === path || (href !== '/' && path.startsWith(href));
    if (isActive) link.classList.add('active');
  });
}
