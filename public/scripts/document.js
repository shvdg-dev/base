// On HTMX events
document.addEventListener('htmx:beforeSwap', (evt) => overruleSwapping(evt));
document.addEventListener('htmx:afterRequest', (evt) => updateTitle(evt));

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
 * Overrule swapping for certain status codes.
 * @param request The request which has the status code.
 */
function overruleSwapping(request){
    const status = request?.detail?.xhr?.status;
    if(status === 401){
        // Allow a 401 response to swap, as it will navigate to a 'authentication required' page.
        request.detail.shouldSwap = true;
    }
}