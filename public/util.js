const html = document.getElementsByTagName('html')[0]

function showContextMenu(elm_type) {
  const contextMenu = document.getElementById('context-menu')
  contextMenu.style.display = "flex"
  html.onclick = () => {
    contextMenu.style.display = "none"
  }
  AppState.set(SELECTED_ENTITY_TYPE, elm_type)
}

function hideContextMenu() {
  const contextMenu = document.getElementById('context-menu')
  contextMenu.style.display = "none"
}

function setTheme(pcolor, scolor, acolor, tcolor) {
  html.style.setProperty("--color-primary", pcolor)
  html.style.setProperty("--color-background", pcolor)
  html.style.setProperty("--color-secondary", scolor)
  html.style.setProperty("--color-accent", acolor)
  html.style.setProperty("--color-text", tcolor)
}

function switchMode() {
  const html = document.getElementsByTagName("html")[0]
  const theme = html.getAttribute("theme")
  html.setAttribute("theme", theme !== "dark" ? "dark" : "light")
}

function blur() {
  const main = document.getElementsByTagName("main")[0]
  main.style.filter = "blur(5px)"
}

function unblur() {
  const main = document.getElementsByTagName("main")[0]
  main.style.filter = ""
}

function closeNotification(target) {
  target.classList.add("animate__fadeOutRight")
  setTimeout(() => target.remove(), 1000);
}

function toggleTruncateClass(elm) {
  if (elm.classList.contains("truncate"))
    elm.classList.remove("truncate")
  else
    elm.classList.add("truncate")
}

bgpattern_elements = []

window.onload = () => {
  const elms = document.querySelectorAll(".bg-pattern")
  const elms2 = document.querySelectorAll(".bg-pattern-v2")
  bgpattern_elements.push(...elms, ...elms2);
}

window.onmousemove = (e) => {
  let size = 60 * e.y/screen.height + 30
  for (elm of bgpattern_elements) {
    elm.style.backgroundSize = `${size}px ${size}px`
    elm.style.backgroundPosition = `${e.pageX}px ${e.pageY}px`
  }
}

// Make elements draggable. credit: https://www.w3schools.com/HOWTO/howto_js_draggable.asp
function dragElement(elmnt) {
  var pos1 = 0, pos2 = 0, pos3 = 0, pos4 = 0;
  const d = document.getElementById(elmnt.id + "-header")
  d.onmousedown = dragMouseDown;
  d.ontouchstart = dragMouseDown;

  function dragMouseDown(e) {
    e.preventDefault();
    // get the mouse cursor position at startup:
    pos3 = e.clientX;
    pos4 = e.clientY;
    document.onmouseup = closeDragElement;
    document.ontouchend = closeDragElement;
    // call a function whenever the cursor moves:
    document.onmousemove = elementDrag;
    document.ontouchmove = elementDrag;
  }

  function elementDrag(e) {
    e.preventDefault();
    // calculate the new cursor position:
    pos1 = pos3 - e.clientX;
    pos2 = pos4 - e.clientY;
    pos3 = e.clientX;
    pos4 = e.clientY;
    // set the element's new position:
    elmnt.style.top = (elmnt.offsetTop - pos2) + "px";
    elmnt.style.left = (elmnt.offsetLeft - pos1) + "px";
  }

  function closeDragElement() {
    // stop moving when mouse button is released:
    document.onmouseup = null;
    document.onmousemove = null;
    document.ontouchend = null;
    document.ontouchmove = null;
  }
}
