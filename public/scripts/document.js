// After a request with HTMX
document.addEventListener('htmx:afterRequest', (evt) => afterRequest(evt), {once: true});

/**
 * Perform tasks after processing a request with HTMX.
 * @param evt The details of the event.
 */
function afterRequest(evt) {
    updateTitle();
    updateMenu();
}

function updateTitle() {

}

function updateMenu(){

}