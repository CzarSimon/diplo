export const DEV_MODE = true;

export const WEBSOCKET_URL = (DEV_MODE)
  ? 'ws://localhost:8080/websocket/chat/connect'
  : 'wss://diploapp.co/websocket/chat/connect'
