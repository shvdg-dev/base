// After loading the DOM
document.addEventListener("DOMContentLoaded", () => refresh(), {once: true});

/**
 * Refreshes the web page if there is no <main> element.
 * Duplicating a tab in the browser fetches the most recent HTML returned from the server.
 * However, with HTMX, these can be snippets (e.g., for swapping).
 * A snippet by itself (no linked files etc.) is not useful, hence, a refresh is done.
 * @return {void} Nothing is returned.
 */
function refresh(){
    if (!document.querySelector("main")) {
        location.reload();
    }
}