!function(){"use strict";const t=["client/client.81f0f175.js","client/en.a7ea2c75.js","client/de.fd2e2b12.js","client/index.490dd4fe.js","client/[slug].bee0947a.js","client/client.0b9139b4.js"].concat(["service-worker-index.html","base.css","fonts/ailerons/ailerons.otf","fonts/fira-mono/fira-mono-latin-400.woff2","fonts/roboto/roboto-latin-400.woff2","fonts/roboto/roboto-latin-400italic.woff2","fonts/roboto/roboto-latin-500.woff2","fonts/roboto/roboto-latin-500italic.woff2","icons/android-chrome-192x192.png","icons/android-chrome-512x512.png","icons/apple-touch-icon.png","icons/favicon-16x16.png","icons/favicon-32x32.png","icons/favicon.ico","icons/mstile-150x150.png","icons/safari-pinned-tab.svg","logos/zitadel-logo-light.svg","logos/zitadel-logo-solo-darkdesign.svg","manifest.json","prism.css"]),e=new Set(t);self.addEventListener("install",e=>{e.waitUntil(caches.open("cache1599222060319").then(e=>e.addAll(t)).then(()=>{self.skipWaiting()}))}),self.addEventListener("activate",t=>{t.waitUntil(caches.keys().then(async t=>{for(const e of t)"cache1599222060319"!==e&&await caches.delete(e);self.clients.claim()}))}),self.addEventListener("fetch",t=>{if("GET"!==t.request.method||t.request.headers.has("range"))return;const o=new URL(t.request.url);o.protocol.startsWith("http")&&(o.hostname===self.location.hostname&&o.port!==self.location.port||(o.host===self.location.host&&e.has(o.pathname)?t.respondWith(caches.match(t.request)):"only-if-cached"!==t.request.cache&&t.respondWith(caches.open("offline1599222060319").then(async e=>{try{const o=await fetch(t.request);return e.put(t.request,o.clone()),o}catch(o){const n=await e.match(t.request);if(n)return n;throw o}}))))})}();
