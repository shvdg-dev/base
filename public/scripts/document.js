// After a request with HTMX
document.addEventListener('htmx:afterRequest', (evt) => updateTitle(evt));
document.addEventListener('htmx:afterRequest', (evt) => updateMenu(evt));

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
 * Update the menu to highlight the selected navigation-item.
 * @param request The request which has the newly selected navigation-item.
 */
function updateMenu(request){
    let path = request?.detail?.pathInfo?.responsePath;
    const menuItems = document.querySelectorAll('[id^="menu-item-"]');
    menuItems.forEach(item => {
        item.classList.add("border-transparent");
        item.classList.remove("border-primary");
        if(path.startsWith(item.id.replace(/menu-item-/g, ''))){
            item.classList.add("border-primary");
        }
    });
}