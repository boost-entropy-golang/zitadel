import{_ as t,a as n,b as e,c as r,i,s as o,d as c,S as a,f as s,o as l,g as u,h as f,v as h,j as v,H as d,k as g,G as m,l as p,D as $,t as w,p as b,w as y,x as _,n as E,y as R,I as z,J as D,K as j,L as x,m as S,M as A,N as B,O as L,P as T,Q as I,r as N,R as k,T as P,q as C,U as V}from"./client.f7b95fec.js";import{_ as q}from"./asyncToGenerator.5229e80b.js";function H(t){var r=function(){if("undefined"==typeof Reflect||!Reflect.construct)return!1;if(Reflect.construct.sham)return!1;if("function"==typeof Proxy)return!0;try{return Date.prototype.toString.call(Reflect.construct(Date,[],(function(){}))),!0}catch(t){return!1}}();return function(){var i,o=n(t);if(r){var c=n(this).constructor;i=Reflect.construct(o,arguments,c)}else i=o.apply(this,arguments);return e(this,i)}}function M(t,n,e){var r=t.slice();return r[12]=n[e],r}function O(t,n,e){var r=t.slice();return r[9]=n[e],r}function U(t){var n,e,r;return e=new L({props:{name:"las la-arrow-right"}}),{c:function(){n=s("div"),b(e.$$.fragment),this.h()},l:function(t){n=u(t,"DIV",{class:!0});var r=f(n);y(e.$$.fragment,r),r.forEach(v),this.h()},h:function(){g(n,"class","icon-container svelte-178w2b1")},m:function(t,i){p(t,n,i),_(e,n,null),r=!0},i:function(t){r||(w(e.$$.fragment,t),r=!0)},o:function(t){E(e.$$.fragment,t),r=!1},d:function(t){t&&v(n),R(e)}}}function G(t){var n,e,r;return e=new L({props:{name:"las la-arrow-right"}}),{c:function(){n=s("div"),b(e.$$.fragment),this.h()},l:function(t){n=u(t,"DIV",{class:!0});var r=f(n);y(e.$$.fragment,r),r.forEach(v),this.h()},h:function(){g(n,"class","icon-container svelte-178w2b1")},m:function(t,i){p(t,n,i),_(e,n,null),r=!0},i:function(t){r||(w(e.$$.fragment,t),r=!0)},o:function(t){E(e.$$.fragment,t),r=!1},d:function(t){t&&v(n),R(e)}}}function J(t){var n,e,r,i,o,c,a=t[12].title+"",b=t[12].slug===t[3]&&G();return{c:function(){n=s("a"),r=l(),b&&b.c(),this.h()},l:function(t){n=u(t,"A",{class:!0,href:!0,"data-level":!0});var e=f(n);r=h(e),b&&b.l(e),e.forEach(v),this.h()},h:function(){e=new d(r),g(n,"class","subsection svelte-178w2b1"),g(n,"href",i=t[1]+"#"+t[12].slug),g(n,"data-level",o=t[12].level),m(n,"active",t[12].slug===t[3])},m:function(t,i){p(t,n,i),e.m(a,n),$(n,r),b&&b.m(n,null),c=!0},p:function(t,r){(!c||4&r)&&a!==(a=t[12].title+"")&&e.p(a),t[12].slug===t[3]?b?12&r&&w(b,1):((b=G()).c(),w(b,1),b.m(n,null)):b&&(z(),E(b,1,1,(function(){b=null})),D()),(!c||6&r&&i!==(i=t[1]+"#"+t[12].slug))&&g(n,"href",i),(!c||4&r&&o!==(o=t[12].level))&&g(n,"data-level",o),12&r&&m(n,"active",t[12].slug===t[3])},i:function(t){c||(w(b),c=!0)},o:function(t){E(b),c=!1},d:function(t){t&&v(n),b&&b.d()}}}function K(t){for(var n,e,r,i,o,c,a,b,y=t[9].metadata.title+"",_=t[9].slug===t[3]&&U(),R=t[9].subsections,x=[],S=0;S<R.length;S+=1)x[S]=J(M(t,R,S));var A=function(t){return E(x[t],1,1,(function(){x[t]=null}))};return{c:function(){n=s("li"),e=s("a"),i=l(),_&&_.c(),c=l();for(var t=0;t<x.length;t+=1)x[t].c();a=l(),this.h()},l:function(t){n=u(t,"LI",{class:!0});var r=f(n);e=u(r,"A",{class:!0,href:!0});var o=f(e);i=h(o),_&&_.l(o),o.forEach(v),c=h(r);for(var s=0;s<x.length;s+=1)x[s].l(r);a=h(r),r.forEach(v),this.h()},h:function(){r=new d(i),g(e,"class","section svelte-178w2b1"),g(e,"href",o=t[1]+"#"+t[9].slug),m(e,"active",t[9].slug===t[3]),g(n,"class","svelte-178w2b1")},m:function(t,o){p(t,n,o),$(n,e),r.m(y,e),$(e,i),_&&_.m(e,null),$(n,c);for(var s=0;s<x.length;s+=1)x[s].m(n,null);$(n,a),b=!0},p:function(t,i){if((!b||4&i)&&y!==(y=t[9].metadata.title+"")&&r.p(y),t[9].slug===t[3]?_?12&i&&w(_,1):((_=U()).c(),w(_,1),_.m(e,null)):_&&(z(),E(_,1,1,(function(){_=null})),D()),(!b||6&i&&o!==(o=t[1]+"#"+t[9].slug))&&g(e,"href",o),12&i&&m(e,"active",t[9].slug===t[3]),14&i){var c;for(R=t[9].subsections,c=0;c<R.length;c+=1){var s=M(t,R,c);x[c]?(x[c].p(s,i),w(x[c],1)):(x[c]=J(s),x[c].c(),w(x[c],1),x[c].m(n,a))}for(z(),c=R.length;c<x.length;c+=1)A(c);D()}},i:function(t){if(!b){w(_);for(var n=0;n<R.length;n+=1)w(x[n]);b=!0}},o:function(t){E(_),x=x.filter(Boolean);for(var n=0;n<x.length;n+=1)E(x[n]);b=!1},d:function(t){t&&v(n),_&&_.d(),j(x,t)}}}function Q(t){for(var n,e,r,i,o=t[2],c=[],a=0;a<o.length;a+=1)c[a]=K(O(t,o,a));var l=function(t){return E(c[t],1,1,(function(){c[t]=null}))};return{c:function(){n=s("ul");for(var t=0;t<c.length;t+=1)c[t].c();this.h()},l:function(t){n=u(t,"UL",{class:!0});for(var e=f(n),r=0;r<c.length;r+=1)c[r].l(e);e.forEach(v),this.h()},h:function(){g(n,"class","reference-toc svelte-178w2b1")},m:function(o,a){p(o,n,a);for(var s=0;s<c.length;s+=1)c[s].m(n,null);t[6](n),e=!0,r||(i=[x(n,"mouseenter",t[7]),x(n,"mouseleave",t[8])],r=!0)},p:function(t,e){var r=S(e,1)[0];if(14&r){var i;for(o=t[2],i=0;i<o.length;i+=1){var a=O(t,o,i);c[i]?(c[i].p(a,r),w(c[i],1)):(c[i]=K(a),c[i].c(),w(c[i],1),c[i].m(n,null))}for(z(),i=o.length;i<c.length;i+=1)l(i);D()}},i:function(t){if(!e){for(var n=0;n<o.length;n+=1)w(c[n]);e=!0}},o:function(t){c=c.filter(Boolean);for(var n=0;n<c.length;n+=1)E(c[n]);e=!1},d:function(e){e&&v(n),j(c,e),t[6](null),r=!1,A(i)}}}function W(t,n,e){var r,i=n.dir,o=void 0===i?"":i,c=n.sections,a=void 0===c?[]:c,s=n.active_section,l=void 0===s?null:s,u=n.show_contents,f=n.prevent_sidebar_scroll,h=void 0!==f&&f;B((function(){if(!(h||u&&window.innerWidth<832)){var t=r.querySelector(".active");if(t){var n=t.getBoundingClientRect(),e=n.top,i=n.bottom,o=window.innerHeight-200;e>o?r.parentNode.scrollBy({top:e-o,left:0,behavior:"smooth"}):i<200&&r.parentNode.scrollBy({top:i-200,left:0,behavior:"smooth"})}}}));return t.$$set=function(t){"dir"in t&&e(1,o=t.dir),"sections"in t&&e(2,a=t.sections),"active_section"in t&&e(3,l=t.active_section),"show_contents"in t&&e(5,u=t.show_contents),"prevent_sidebar_scroll"in t&&e(0,h=t.prevent_sidebar_scroll)},[h,o,a,l,r,u,function(t){T[t?"unshift":"push"]((function(){e(4,r=t)}))},function(){return e(0,h=!0)},function(){return e(0,h=!1)}]}var Y=function(n){t(s,a);var e=H(s);function s(t){var n;return r(this,s),n=e.call(this),i(c(n),t,W,Q,o,{dir:1,sections:2,active_section:3,show_contents:5,prevent_sidebar_scroll:0}),n}return s}();function F(t){var r=function(){if("undefined"==typeof Reflect||!Reflect.construct)return!1;if(Reflect.construct.sham)return!1;if("function"==typeof Proxy)return!0;try{return Date.prototype.toString.call(Reflect.construct(Date,[],(function(){}))),!0}catch(t){return!1}}();return function(){var i,o=n(t);if(r){var c=n(this).constructor;i=Reflect.construct(o,arguments,c)}else i=o.apply(this,arguments);return e(this,i)}}function X(t,n,e){var r=t.slice();return r[14]=n[e],r}function Z(t){var n,e,r,i,o,c,a,m,z,D,j,x,S,A,B,T,I,N,k,P=t[14].metadata.title+"",C=t[14].html+"";return S=new L({props:{name:"las la-external-link-alt",size:"24px"}}),{c:function(){n=s("section"),e=s("h2"),r=s("span"),o=l(),c=s("a"),m=l(),D=l(),j=s("small"),x=s("a"),b(S.$$.fragment),B=l(),I=l(),this.h()},l:function(t){n=u(t,"SECTION",{"data-id":!0,class:!0});var i=f(n);e=u(i,"H2",{class:!0});var a=f(e);r=u(a,"SPAN",{class:!0,id:!0}),f(r).forEach(v),o=h(a),c=u(a,"A",{href:!0,class:!0,"aria-hidden":!0}),f(c).forEach(v),m=h(a),D=h(a),j=u(a,"SMALL",{class:!0});var s=f(j);x=u(s,"A",{href:!0,title:!0,class:!0});var l=f(x);y(S.$$.fragment,l),l.forEach(v),s.forEach(v),a.forEach(v),B=h(i),I=h(i),i.forEach(v),this.h()},h:function(){g(r,"class","offset-anchor"),g(r,"id",i=t[14].slug),g(c,"href",a=t[3]+"#"+t[14].slug),g(c,"class","anchor"),g(c,"aria-hidden",""),z=new d(D),g(x,"href",A="https://github.com/"+t[0]+"/"+t[1]+"/edit/master"+t[2]+"/"+t[3]+"/"+t[14].file),g(x,"title",t[4]),g(x,"class","svelte-eu3cmz"),g(j,"class","svelte-eu3cmz"),g(e,"class","svelte-eu3cmz"),T=new d(I),g(n,"data-id",N=t[14].slug),g(n,"class","svelte-eu3cmz")},m:function(t,i){p(t,n,i),$(n,e),$(e,r),$(e,o),$(e,c),$(e,m),z.m(P,e),$(e,D),$(e,j),$(j,x),_(S,x,null),$(n,B),T.m(C,n),$(n,I),k=!0},p:function(t,e){(!k||32&e&&i!==(i=t[14].slug))&&g(r,"id",i),(!k||40&e&&a!==(a=t[3]+"#"+t[14].slug))&&g(c,"href",a),(!k||32&e)&&P!==(P=t[14].metadata.title+"")&&z.p(P),(!k||47&e&&A!==(A="https://github.com/"+t[0]+"/"+t[1]+"/edit/master"+t[2]+"/"+t[3]+"/"+t[14].file))&&g(x,"href",A),(!k||16&e)&&g(x,"title",t[4]),(!k||32&e)&&C!==(C=t[14].html+"")&&T.p(C),(!k||32&e&&N!==(N=t[14].slug))&&g(n,"data-id",N)},i:function(t){k||(w(S.$$.fragment,t),k=!0)},o:function(t){E(S.$$.fragment,t),k=!1},d:function(t){t&&v(n),R(S)}}}function tt(t){for(var n,e,r,i,o,c,a,d,B,T,I,N=t[5],k=[],P=0;P<N.length;P+=1)k[P]=Z(X(t,N,P));var C=function(t){return E(k[t],1,1,(function(){k[t]=null}))};return o=new Y({props:{dir:t[3],sections:t[5],active_section:t[6],show_contents:t[9]}}),d=new L({props:{name:t[9]?"las la-window-close":"las la-bars"}}),{c:function(){n=s("div");for(var t=0;t<k.length;t+=1)k[t].c();e=l(),r=s("aside"),i=s("div"),b(o.$$.fragment),c=l(),a=s("button"),b(d.$$.fragment),this.h()},l:function(t){n=u(t,"DIV",{class:!0});for(var s=f(n),l=0;l<k.length;l+=1)k[l].l(s);s.forEach(v),e=h(t),r=u(t,"ASIDE",{class:!0});var g=f(r);i=u(g,"DIV",{class:!0});var m=f(i);y(o.$$.fragment,m),m.forEach(v),c=h(g),a=u(g,"BUTTON",{class:!0});var p=f(a);y(d.$$.fragment,p),p.forEach(v),g.forEach(v),this.h()},h:function(){g(n,"class","content listify svelte-eu3cmz"),g(i,"class","sidebar svelte-eu3cmz"),g(a,"class","svelte-eu3cmz"),g(r,"class","sidebar-container svelte-eu3cmz"),m(r,"open",t[9])},m:function(s,l){p(s,n,l);for(var u=0;u<k.length;u+=1)k[u].m(n,null);t[10](n),p(s,e,l),p(s,r,l),$(r,i),_(o,i,null),$(r,c),$(r,a),_(d,a,null),t[13](r),B=!0,T||(I=[x(i,"click",t[11]),x(a,"click",t[12])],T=!0)},p:function(t,e){var i=S(e,1)[0];if(63&i){var c;for(N=t[5],c=0;c<N.length;c+=1){var a=X(t,N,c);k[c]?(k[c].p(a,i),w(k[c],1)):(k[c]=Z(a),k[c].c(),w(k[c],1),k[c].m(n,null))}for(z(),c=N.length;c<k.length;c+=1)C(c);D()}var s={};8&i&&(s.dir=t[3]),32&i&&(s.sections=t[5]),64&i&&(s.active_section=t[6]),512&i&&(s.show_contents=t[9]),o.$set(s);var l={};512&i&&(l.name=t[9]?"las la-window-close":"las la-bars"),d.$set(l),512&i&&m(r,"open",t[9])},i:function(t){if(!B){for(var n=0;n<N.length;n+=1)w(k[n]);w(o.$$.fragment,t),w(d.$$.fragment,t),B=!0}},o:function(t){k=k.filter(Boolean);for(var n=0;n<k.length;n+=1)E(k[n]);E(o.$$.fragment,t),E(d.$$.fragment,t),B=!1},d:function(i){i&&v(n),j(k,i),t[10](null),i&&v(e),i&&v(r),R(o),R(d),t[13](null),T=!1,A(I)}}}function nt(t,n,e){var r,i,o,c=n.owner,a=void 0===c?"caos":c,s=n.project,l=void 0===s?"zitadel":s,u=n.path,f=void 0===u?"site/docs":u,h=n.dir,v=void 0===h?"":h,d=n.edit_title,g=void 0===d?"edit this section":d,m=n.sections,p=!1;I((function(){var t,n=i.querySelectorAll("[id]:not([data-scrollignore])"),o=function(){var e=i.getBoundingClientRect().top;t=[].map.call(n,(function(t){return t.getBoundingClientRect().top-e}))},c=window.location.hash.slice(1),a=function(){for(var i=-window.scrollY,o=n.length;o--;)if(t[o]+i<40){var a=n[o].id;return void(a!==c&&(e(6,r=a),c=a))}};window.addEventListener("scroll",a,!0),window.addEventListener("resize",o,!0);var s=[setTimeout(o,1e3),setTimeout(a,5e3)];return o(),a(),function(){window.removeEventListener("scroll",a,!0),window.removeEventListener("resize",o,!0),s.forEach((function(t){return clearTimeout(t)}))}}));return t.$$set=function(t){"owner"in t&&e(0,a=t.owner),"project"in t&&e(1,l=t.project),"path"in t&&e(2,f=t.path),"dir"in t&&e(3,v=t.dir),"edit_title"in t&&e(4,g=t.edit_title),"sections"in t&&e(5,m=t.sections)},[a,l,f,v,g,m,r,i,o,p,function(t){T[t?"unshift":"push"]((function(){e(7,i=t)}))},function(){return e(9,p=!1)},function(){return e(9,p=!p)},function(t){T[t?"unshift":"push"]((function(){e(8,o=t)}))}]}var et=function(n){t(s,a);var e=F(s);function s(t){var n;return r(this,s),n=e.call(this),i(c(n),t,nt,tt,o,{owner:0,project:1,path:2,dir:3,edit_title:4,sections:5}),n}return s}();function rt(t){var r=function(){if("undefined"==typeof Reflect||!Reflect.construct)return!1;if(Reflect.construct.sham)return!1;if("function"==typeof Proxy)return!0;try{return Date.prototype.toString.call(Reflect.construct(Date,[],(function(){}))),!0}catch(t){return!1}}();return function(){var i,o=n(t);if(r){var c=n(this).constructor;i=Reflect.construct(o,arguments,c)}else i=o.apply(this,arguments);return e(this,i)}}function it(t,n,e){var r=t.slice();return r[3]=n[e].name,r[4]=n[e].content,r[6]=e,r}function ot(t){var n,e,r;return{c:function(){n=s("meta"),this.h()},l:function(t){n=u(t,"META",{name:!0,content:!0}),this.h()},h:function(){g(n,"name",e=t[3]),g(n,"content",r=t[4])},m:function(t,e){p(t,n,e)},p:function(t,i){4&i&&e!==(e=t[3])&&g(n,"name",e),4&i&&r!==(r=t[4])&&g(n,"content",r)},d:function(t){t&&v(n)}}}function ct(t){var n,e,r,i,o;document.title=n=k.name+" • "+t[0];for(var c=t[2],a=[],s=0;s<c.length;s+=1)a[s]=ot(it(t,c,s));return i=new et({props:{sections:t[1],project:"zitadel/site",dir:t[0]}}),{c:function(){for(var t=0;t<a.length;t+=1)a[t].c();e=P(),r=l(),b(i.$$.fragment)},l:function(t){for(var n=C('[data-svelte="svelte-4bltt5"]',document.head),o=0;o<a.length;o+=1)a[o].l(n);e=P(),n.forEach(v),r=h(t),y(i.$$.fragment,t)},m:function(t,n){for(var c=0;c<a.length;c+=1)a[c].m(document.head,null);$(document.head,e),p(t,r,n),_(i,t,n),o=!0},p:function(t,r){var s=S(r,1)[0];if((!o||1&s)&&n!==(n=k.name+" • "+t[0])&&(document.title=n),4&s){var l;for(c=t[2],l=0;l<c.length;l+=1){var u=it(t,c,l);a[l]?a[l].p(u,s):(a[l]=ot(u),a[l].c(),a[l].m(e.parentNode,e))}for(;l<a.length;l+=1)a[l].d(1);a.length=c.length}var f={};2&s&&(f.sections=t[1]),1&s&&(f.dir=t[0]),i.$set(f)},i:function(t){o||(w(i.$$.fragment,t),o=!0)},o:function(t){E(i.$$.fragment,t),o=!1},d:function(t){j(a,t),v(e),t&&v(r),R(i,t)}}}function at(t){return st.apply(this,arguments)}function st(){return(st=q(N.mark((function t(n){var e,r,i,o;return N.wrap((function(t){for(;;)switch(t.prev=t.next){case 0:return e=n.params,e.lang,r=e.slug,t.next=4,this.fetch("".concat(r,".json")).then((function(t){return t.json()}));case 4:return i=t.sent,o=[],t.abrupt("return",{sections:i,slug:r,tags:o});case 7:case"end":return t.stop()}}),t,this)})))).apply(this,arguments)}function lt(t,n,e){var r=n.slug,i=n.sections,o=n.tags;return I((function(){V(".zitadel-gallery")})),t.$$set=function(t){"slug"in t&&e(0,r=t.slug),"sections"in t&&e(1,i=t.sections),"tags"in t&&e(2,o=t.tags)},[r,i,o]}var ut=function(n){t(s,a);var e=rt(s);function s(t){var n;return r(this,s),n=e.call(this),i(c(n),t,lt,ct,o,{slug:0,sections:1,tags:2}),n}return s}();export default ut;export{at as preload};
