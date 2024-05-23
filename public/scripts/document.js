// On HTMX events
document.addEventListener('htmx:beforeSwap', (evt) => overruleSwapping(evt));
document.addEventListener('htmx:afterRequest', (evt) => updateTitle(evt));
document.addEventListener('htmx:afterRequest', (evt) => updateURL(evt));

/**
 * Update the title, a.k.a. the tab name.
 * @param request The request which has the new title.
 */
function updateTitle(request) {
    let newTitle = request.detail.xhr.getResponseHeader('H-Title');
    if (newTitle) {
        document.title = newTitle;
    }
}

/**
 * Update the URL in the address bar, this does not visit the URL, only displays it.
 * @param request The request which has the new path.
 */
function updateURL(request) {
    let newPath = request.detail.xhr.getResponseHeader('H-Path');
    if (newPath) {
        history.pushState(null, null, newPath);
    }
}


/**
 * Overrule swapping for certain status codes.
 * @param request The request which has the status code.
 */
function overruleSwapping(request){
    const status = request?.detail?.xhr?.status;
    if(status === 401 || status === 404){
        // Allow a 4** response to swap, as it will navigate to the corresponding error page.
        request.detail.shouldSwap = true;
    }
}