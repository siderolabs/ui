// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

function getID(contextName, runtimeObject) {
  return contextName + "/" + runtimeObject;
}

function EventListener() {
  this.subscriptions = {};
  window.wails.Events.On("capiEvent", this.handleCAPIEvent.bind(this))
}

EventListener.prototype.bind = function(contextName, runtimeObject, handler) {
  var id = getID(contextName, runtimeObject);
  if(!this.subscriptions[id]) {
    this.subscriptions[id] = [];
  }

  this.subscriptions[id].push(handler);
  return window.backend.CAPI.Subscribe(contextName, runtimeObject);
}

EventListener.prototype.unbind = function(contextName, runtimeObject, handler) {
  var id = getID(contextName, runtimeObject);
  var handlers = this.subscriptions[id];

  if(!handlers) {
    return;
  }

  for(var i = 0; i < handlers.length; i++) {
    if(handlers[i] === handler) {
      handlers.splice(i, 1);
    }
  }

  if(handlers.length == 0) {
    delete this.subscriptions[id];
    return window.backend.CAPI.Unsubscribe(contextName, runtimeObject);
  }

  return Promise.resolve();
}

EventListener.prototype.handleCAPIEvent = function(e) {
  var id = getID(e.contextName, e.runtimeObject);
  if(this.subscriptions[id]) {
    this.subscriptions[id].forEach(handler => handler(e));
  }
}

var eventListener = new EventListener();

export default {

  methods: {
    bind(contextName, runtimeObject, handler) {
      return eventListener.bind(contextName, runtimeObject, handler);
    },

    unbind(contextName, runtimeObject, handler) {
      return eventListener.unbind(contextName, runtimeObject, handler);
    },
  }
};
