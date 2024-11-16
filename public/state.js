// some constants used as state values keys for consistency
// and for reflecting the application concept in code 
const SELECTED_CHANNEL = "sc";
const SELECTED_SECTION = "ss";
const SELECTED_VIDEO = "sv";

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

console.log(
  SELECTED_CHANNEL,
  SELECTED_SECTION,
  SELECTED_VIDEO,
  AppState
)

// explicitly set listeners
AppState.addListener(SELECTED_CHANNEL, () => {
  const btns = document.getElementsByClassName("channel-btn");
  for (const btn of btns) {
    if (btn.dataset.index == AppState.get(SELECTED_CHANNEL)) {
      btn.style.opacity = "1";
    }
    else {
      btn.style.opacity = "0.65";
    }
  }
})
