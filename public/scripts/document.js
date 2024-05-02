document.addEventListener("DOMContentLoaded", () => {
    if (!document.querySelector("main")) {
        htmx.ajax('GET', window.location.href, {swap: 'body.innerHTML'});
    }
}, {once: true});