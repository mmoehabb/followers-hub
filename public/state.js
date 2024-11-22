// some constants used as state values keys for consistency
// and for reflecting the application concept in code 
const SELECTED_CHANNEL = "SC";
const SELECTED_SECTION = "SS";
const SELECTED_VIDEO = "SV";
const SELECTED_ENTITY_TYPE = "SET" // used in context menu

const TYPE_CHANNEL = "channel"
const TYPE_SECTION = "section"
const TYPE_VIDEO = "video"

class StateManager {
  _stateMap = {};
  _listeners = {};

  set(key, value) {
    this._stateMap[key] = value;
    if (this._listeners[key]) {
      for (const func of this._listeners[key]) func();
    }
  }

  get(key) {
    return this._stateMap[key];
  }

  addListener(key, func) {
    if (!this._listeners[key]) {
      this._listeners[key] = [];
    }
    this._listeners[key].push(func);
  }
}

const AppState = new StateManager();

// explicitly set listeners:
// for channels buttons
AppState.addListener(SELECTED_CHANNEL, () => {
  const btns = document.getElementsByClassName("channel-btn");
  for (const btn of btns) {
    if (btn.dataset.id == AppState.get(SELECTED_CHANNEL)) {
      btn.style.opacity = "1";
    }
    else {
      btn.style.opacity = "0.65";
    }
  }
})
// for sections buttons
AppState.addListener(SELECTED_SECTION, () => {
  const btns = document.getElementsByClassName("section-btn");
  for (const btn of btns) {
    if (btn.dataset.id == AppState.get(SELECTED_SECTION)) {
      btn.style.opacity = "1";
    }
    else {
      btn.style.opacity = "0.65";
    }
  }
})
// for the context menu
AppState.addListener(SELECTED_ENTITY_TYPE, () => {
  const ctx_menu_label = document.getElementById("context-menu-label");
  ctx_menu_label.innerText = "Options for selected " + AppState.get(SELECTED_ENTITY_TYPE);
})

console.log(
  SELECTED_CHANNEL,
  SELECTED_SECTION,
  SELECTED_VIDEO,
  SELECTED_ENTITY_TYPE,

  TYPE_VIDEO,
  TYPE_CHANNEL,
  TYPE_SECTION,

  AppState
)

