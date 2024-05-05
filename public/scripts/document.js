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
 * Update the menu to highlight the selected menu-item.
 * @param request The request for determining highlighting.
 */
function updateMenu(request) {
    const path = request?.detail?.pathInfo?.responsePath;
    const menuItemPrefix = 'menu-item-';
    const menuItems = document.querySelectorAll(`[id^="${menuItemPrefix}"]`);

    menuItems.forEach(modifyMenuItem);

    function modifyMenuItem(menuItem) {
        menuItem.classList.add("border-transparent");

        const menuItemIdWithoutPrefix = menuItem.id.replace(new RegExp(menuItemPrefix, 'g'), '');
        const shouldHighlightItem = path !== null && path.startsWith(menuItemIdWithoutPrefix);

        if (shouldHighlightItem) {
            menuItem.classList.remove("border-transparent");
            menuItem.classList.add("border-primary");
        }
    }
}