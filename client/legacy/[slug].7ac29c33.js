import{_ as n,a as t,b as e,c as r,i,s as o,d as c,S as a,f as s,p as l,g as u,h as f,w as h,j as v,H as d,k as g,G as m,l as p,D as $,t as w,q as y,x as E,y as _,n as b,z as R,I as D,J as x,K as S,L as j,m as A,M as B,N as L,O as I,P as T,Q as N,R as z,o as k,r as P,T as C,v as V}from"./client.74d3c8ff.js";function q(n){var r=function(){if("undefined"==typeof Reflect||!Reflect.construct)return!1;if(Reflect.construct.sham)return!1;if("function"==typeof Proxy)return!0;try{return Date.prototype.toString.call(Reflect.construct(Date,[],(function(){}))),!0}catch(n){return!1}}();return function(){var i,o=t(n);if(r){var c=t(this).constructor;i=Reflect.construct(o,arguments,c)}else i=o.apply(this,arguments);return e(this,i)}}function H(n,t,e){var r=n.slice();return r[12]=t[e],r}function M(n,t,e){var r=n.slice();return r[9]=t[e],r}function O(n){var t,e,r;return e=new I({props:{name:"las la-arrow-right"}}),{c:function(){t=s("div"),y(e.$$.fragment),this.h()},l:function(n){t=u(n,"DIV",{class:!0});var r=f(t);E(e.$$.fragment,r),r.forEach(v),this.h()},h:function(){g(t,"class","icon-container svelte-1s4ps85")},m:function(n,i){p(n,t,i),_(e,t,null),r=!0},i:function(n){r||(w(e.$$.fragment,n),r=!0)},o:function(n){b(e.$$.fragment,n),r=!1},d:function(n){n&&v(t),R(e)}}}function U(n){var t,e,r;return e=new I({props:{name:"las la-arrow-right"}}),{c:function(){t=s("div"),y(e.$$.fragment),this.h()},l:function(n){t=u(n,"DIV",{class:!0});var r=f(t);E(e.$$.fragment,r),r.forEach(v),this.h()},h:function(){g(t,"class","icon-container svelte-1s4ps85")},m:function(n,i){p(n,t,i),_(e,t,null),r=!0},i:function(n){r||(w(e.$$.fragment,n),r=!0)},o:function(n){b(e.$$.fragment,n),r=!1},d:function(n){n&&v(t),R(e)}}}function G(n){var t,e,r,i,o,c,a=n[12].title+"",y=n[12].slug===n[3]&&U();return{c:function(){t=s("a"),r=l(),y&&y.c(),this.h()},l:function(n){t=u(n,"A",{class:!0,href:!0,"data-level":!0});var e=f(t);r=h(e),y&&y.l(e),e.forEach(v),this.h()},h:function(){e=new d(r),g(t,"class","subsection svelte-1s4ps85"),g(t,"href",i=n[1]+"#"+n[12].slug),g(t,"data-level",o=n[12].level),m(t,"active",n[12].slug===n[3])},m:function(n,i){p(n,t,i),e.m(a,t),$(t,r),y&&y.m(t,null),c=!0},p:function(n,r){(!c||4&r)&&a!==(a=n[12].title+"")&&e.p(a),n[12].slug===n[3]?y?12&r&&w(y,1):((y=U()).c(),w(y,1),y.m(t,null)):y&&(D(),b(y,1,1,(function(){y=null})),x()),(!c||6&r&&i!==(i=n[1]+"#"+n[12].slug))&&g(t,"href",i),(!c||4&r&&o!==(o=n[12].level))&&g(t,"data-level",o),12&r&&m(t,"active",n[12].slug===n[3])},i:function(n){c||(w(y),c=!0)},o:function(n){b(y),c=!1},d:function(n){n&&v(t),y&&y.d()}}}function J(n){for(var t,e,r,i,o,c,a,y,E=n[9].metadata.title+"",_=n[9].slug===n[3]&&O(),R=n[9].subsections,j=[],A=0;A<R.length;A+=1)j[A]=G(H(n,R,A));var B=function(n){return b(j[n],1,1,(function(){j[n]=null}))};return{c:function(){t=s("li"),e=s("a"),i=l(),_&&_.c(),c=l();for(var n=0;n<j.length;n+=1)j[n].c();a=l(),this.h()},l:function(n){t=u(n,"LI",{class:!0});var r=f(t);e=u(r,"A",{class:!0,href:!0});var o=f(e);i=h(o),_&&_.l(o),o.forEach(v),c=h(r);for(var s=0;s<j.length;s+=1)j[s].l(r);a=h(r),r.forEach(v),this.h()},h:function(){r=new d(i),g(e,"class","section svelte-1s4ps85"),g(e,"href",o=n[1]+"#"+n[9].slug),m(e,"active",n[9].slug===n[3]),g(t,"class","svelte-1s4ps85")},m:function(n,o){p(n,t,o),$(t,e),r.m(E,e),$(e,i),_&&_.m(e,null),$(t,c);for(var s=0;s<j.length;s+=1)j[s].m(t,null);$(t,a),y=!0},p:function(n,i){if((!y||4&i)&&E!==(E=n[9].metadata.title+"")&&r.p(E),n[9].slug===n[3]?_?12&i&&w(_,1):((_=O()).c(),w(_,1),_.m(e,null)):_&&(D(),b(_,1,1,(function(){_=null})),x()),(!y||6&i&&o!==(o=n[1]+"#"+n[9].slug))&&g(e,"href",o),12&i&&m(e,"active",n[9].slug===n[3]),14&i){var c;for(R=n[9].subsections,c=0;c<R.length;c+=1){var s=H(n,R,c);j[c]?(j[c].p(s,i),w(j[c],1)):(j[c]=G(s),j[c].c(),w(j[c],1),j[c].m(t,a))}for(D(),c=R.length;c<j.length;c+=1)B(c);x()}},i:function(n){if(!y){w(_);for(var t=0;t<R.length;t+=1)w(j[t]);y=!0}},o:function(n){b(_),j=j.filter(Boolean);for(var t=0;t<j.length;t+=1)b(j[t]);y=!1},d:function(n){n&&v(t),_&&_.d(),S(j,n)}}}function K(n){for(var t,e,r,i,o=n[2],c=[],a=0;a<o.length;a+=1)c[a]=J(M(n,o,a));var l=function(n){return b(c[n],1,1,(function(){c[n]=null}))};return{c:function(){t=s("ul");for(var n=0;n<c.length;n+=1)c[n].c();this.h()},l:function(n){t=u(n,"UL",{class:!0});for(var e=f(t),r=0;r<c.length;r+=1)c[r].l(e);e.forEach(v),this.h()},h:function(){g(t,"class","reference-toc svelte-1s4ps85")},m:function(o,a){p(o,t,a);for(var s=0;s<c.length;s+=1)c[s].m(t,null);n[6](t),e=!0,r||(i=[j(t,"mouseenter",n[7]),j(t,"mouseleave",n[8])],r=!0)},p:function(n,e){var r=A(e,1)[0];if(14&r){var i;for(o=n[2],i=0;i<o.length;i+=1){var a=M(n,o,i);c[i]?(c[i].p(a,r),w(c[i],1)):(c[i]=J(a),c[i].c(),w(c[i],1),c[i].m(t,null))}for(D(),i=o.length;i<c.length;i+=1)l(i);x()}},i:function(n){if(!e){for(var t=0;t<o.length;t+=1)w(c[t]);e=!0}},o:function(n){c=c.filter(Boolean);for(var t=0;t<c.length;t+=1)b(c[t]);e=!1},d:function(e){e&&v(t),S(c,e),n[6](null),r=!1,B(i)}}}function Q(n,t,e){var r,i=t.dir,o=void 0===i?"":i,c=t.sections,a=void 0===c?[]:c,s=t.active_section,l=void 0===s?null:s,u=t.show_contents,f=t.prevent_sidebar_scroll,h=void 0!==f&&f;L((function(){if(!(h||u&&window.innerWidth<832)){var n=r.querySelector(".active");if(n){var t=n.getBoundingClientRect(),e=t.top,i=t.bottom,o=window.innerHeight-200;e>o?r.parentNode.scrollBy({top:e-o,left:0,behavior:"smooth"}):i<200&&r.parentNode.scrollBy({top:i-200,left:0,behavior:"smooth"})}}}));return n.$set=function(n){"dir"in n&&e(1,o=n.dir),"sections"in n&&e(2,a=n.sections),"active_section"in n&&e(3,l=n.active_section),"show_contents"in n&&e(5,u=n.show_contents),"prevent_sidebar_scroll"in n&&e(0,h=n.prevent_sidebar_scroll)},[h,o,a,l,r,u,function(n){T[n?"unshift":"push"]((function(){e(4,r=n)}))},function(){return e(0,h=!0)},function(){return e(0,h=!1)}]}var W=function(t){n(s,a);var e=q(s);function s(n){var t;return r(this,s),t=e.call(this),i(c(t),n,Q,K,o,{dir:1,sections:2,active_section:3,show_contents:5,prevent_sidebar_scroll:0}),t}return s}();function Y(n){var r=function(){if("undefined"==typeof Reflect||!Reflect.construct)return!1;if(Reflect.construct.sham)return!1;if("function"==typeof Proxy)return!0;try{return Date.prototype.toString.call(Reflect.construct(Date,[],(function(){}))),!0}catch(n){return!1}}();return function(){var i,o=t(n);if(r){var c=t(this).constructor;i=Reflect.construct(o,arguments,c)}else i=o.apply(this,arguments);return e(this,i)}}function F(n,t,e){var r=n.slice();return r[15]=t[e],r}function X(n){var t,e,r,i,o,c,a,m,D,x,S,j,A,B,L,T,N,z,k,P=n[15].metadata.title+"",C=n[15].html+"";return A=new I({props:{name:"las la-external-link-alt",size:"24px"}}),{c:function(){t=s("section"),e=s("h2"),r=s("span"),o=l(),c=s("a"),m=l(),x=l(),S=s("small"),j=s("a"),y(A.$$.fragment),L=l(),N=l(),this.h()},l:function(n){t=u(n,"SECTION",{"data-id":!0,class:!0});var i=f(t);e=u(i,"H2",{class:!0});var a=f(e);r=u(a,"SPAN",{class:!0,id:!0}),f(r).forEach(v),o=h(a),c=u(a,"A",{href:!0,class:!0,"aria-hidden":!0}),f(c).forEach(v),m=h(a),x=h(a),S=u(a,"SMALL",{class:!0});var s=f(S);j=u(s,"A",{href:!0,title:!0,class:!0});var l=f(j);E(A.$$.fragment,l),l.forEach(v),s.forEach(v),a.forEach(v),L=h(i),N=h(i),i.forEach(v),this.h()},h:function(){g(r,"class","offset-anchor"),g(r,"id",i=n[15].slug),g(c,"href",a=n[3]+"#"+n[15].slug),g(c,"class","anchor"),g(c,"aria-hidden",""),D=new d(x),g(j,"href",B="https://github.com/"+n[0]+"/"+n[1]+"/edit/master"+n[2]+"/"+n[3]+"/"+n[15].file),g(j,"title",n[4]),g(j,"class","svelte-5y8v90"),g(S,"class","svelte-5y8v90"),g(e,"class","svelte-5y8v90"),T=new d(N),g(t,"data-id",z=n[15].slug),g(t,"class","svelte-5y8v90")},m:function(n,i){p(n,t,i),$(t,e),$(e,r),$(e,o),$(e,c),$(e,m),D.m(P,e),$(e,x),$(e,S),$(S,j),_(A,j,null),$(t,L),T.m(C,t),$(t,N),k=!0},p:function(n,e){(!k||32&e&&i!==(i=n[15].slug))&&g(r,"id",i),(!k||40&e&&a!==(a=n[3]+"#"+n[15].slug))&&g(c,"href",a),(!k||32&e)&&P!==(P=n[15].metadata.title+"")&&D.p(P),(!k||47&e&&B!==(B="https://github.com/"+n[0]+"/"+n[1]+"/edit/master"+n[2]+"/"+n[3]+"/"+n[15].file))&&g(j,"href",B),(!k||16&e)&&g(j,"title",n[4]),(!k||32&e)&&C!==(C=n[15].html+"")&&T.p(C),(!k||32&e&&z!==(z=n[15].slug))&&g(t,"data-id",z)},i:function(n){k||(w(A.$$.fragment,n),k=!0)},o:function(n){b(A.$$.fragment,n),k=!1},d:function(n){n&&v(t),R(A)}}}function Z(n){for(var t,e,r,i,o,c,a,d,L,T,N,z=n[5],k=[],P=0;P<z.length;P+=1)k[P]=X(F(n,z,P));var C=function(n){return b(k[n],1,1,(function(){k[n]=null}))};return o=new W({props:{dir:n[3],sections:n[5],active_section:n[6],show_contents:n[9]}}),d=new I({props:{name:n[9]?"las la-window-close":"las la-bars"}}),{c:function(){t=s("div");for(var n=0;n<k.length;n+=1)k[n].c();e=l(),r=s("aside"),i=s("div"),y(o.$$.fragment),c=l(),a=s("button"),y(d.$$.fragment),this.h()},l:function(n){t=u(n,"DIV",{class:!0});for(var s=f(t),l=0;l<k.length;l+=1)k[l].l(s);s.forEach(v),e=h(n),r=u(n,"ASIDE",{class:!0});var g=f(r);i=u(g,"DIV",{class:!0});var m=f(i);E(o.$$.fragment,m),m.forEach(v),c=h(g),a=u(g,"BUTTON",{class:!0});var p=f(a);E(d.$$.fragment,p),p.forEach(v),g.forEach(v),this.h()},h:function(){g(t,"class","content listify svelte-5y8v90"),g(i,"class","sidebar svelte-5y8v90"),g(a,"class","svelte-5y8v90"),g(r,"class","sidebar-container svelte-5y8v90"),m(r,"open",n[9])},m:function(s,l){p(s,t,l);for(var u=0;u<k.length;u+=1)k[u].m(t,null);n[11](t),p(s,e,l),p(s,r,l),$(r,i),_(o,i,null),$(r,c),$(r,a),_(d,a,null),n[14](r),L=!0,T||(N=[j(i,"click",n[12]),j(a,"click",n[13])],T=!0)},p:function(n,e){var i=A(e,1)[0];if(63&i){var c;for(z=n[5],c=0;c<z.length;c+=1){var a=F(n,z,c);k[c]?(k[c].p(a,i),w(k[c],1)):(k[c]=X(a),k[c].c(),w(k[c],1),k[c].m(t,null))}for(D(),c=z.length;c<k.length;c+=1)C(c);x()}var s={};8&i&&(s.dir=n[3]),32&i&&(s.sections=n[5]),64&i&&(s.active_section=n[6]),512&i&&(s.show_contents=n[9]),o.$set(s);var l={};512&i&&(l.name=n[9]?"las la-window-close":"las la-bars"),d.$set(l),512&i&&m(r,"open",n[9])},i:function(n){if(!L){for(var t=0;t<z.length;t+=1)w(k[t]);w(o.$$.fragment,n),w(d.$$.fragment,n),L=!0}},o:function(n){k=k.filter(Boolean);for(var t=0;t<k.length;t+=1)b(k[t]);b(o.$$.fragment,n),b(d.$$.fragment,n),L=!1},d:function(i){i&&v(t),S(k,i),n[11](null),i&&v(e),i&&v(r),R(o),R(d),n[14](null),T=!1,B(N)}}}function nn(n,t,e){var r,i,o,c=t.lang,a=void 0===c?z.lang:c,s=t.owner,l=void 0===s?"caos":s,u=t.project,f=void 0===u?"zitadel/site":u,h=t.path,v=void 0===h?"/docs":h,d=t.dir,g=void 0===d?"":d,m=t.edit_title,p=void 0===m?"edit this section":m,$=t.sections,w=!1;N((function(){var n,t=i.querySelectorAll("[id]:not([data-scrollignore])"),o=function(){var e=i.getBoundingClientRect().top;n=[].map.call(t,(function(n){return n.getBoundingClientRect().top-e}))},c=window.location.hash.slice(1),a=function(){for(var i=-window.scrollY,o=t.length;o--;)if(n[o]+i<40){var a=t[o].id;return void(a!==c&&(e(6,r=a),c=a))}};window.addEventListener("scroll",a,!0),window.addEventListener("resize",o,!0);var s=[setTimeout(o,1e3),setTimeout(a,5e3)];return o(),a(),function(){window.removeEventListener("scroll",a,!0),window.removeEventListener("resize",o,!0),s.forEach((function(n){return clearTimeout(n)}))}}));return n.$set=function(n){"lang"in n&&e(10,a=n.lang),"owner"in n&&e(0,l=n.owner),"project"in n&&e(1,f=n.project),"path"in n&&e(2,v=n.path),"dir"in n&&e(3,g=n.dir),"edit_title"in n&&e(4,p=n.edit_title),"sections"in n&&e(5,$=n.sections)},[l,f,v,g,p,$,r,i,o,w,a,function(n){T[n?"unshift":"push"]((function(){e(7,i=n)}))},function(){return e(9,w=!1)},function(){return e(9,w=!w)},function(n){T[n?"unshift":"push"]((function(){e(8,o=n)}))}]}var tn=function(t){n(s,a);var e=Y(s);function s(n){var t;return r(this,s),t=e.call(this),i(c(t),n,nn,Z,o,{lang:10,owner:0,project:1,path:2,dir:3,edit_title:4,sections:5}),t}return s}();function en(n){var r=function(){if("undefined"==typeof Reflect||!Reflect.construct)return!1;if(Reflect.construct.sham)return!1;if("function"==typeof Proxy)return!0;try{return Date.prototype.toString.call(Reflect.construct(Date,[],(function(){}))),!0}catch(n){return!1}}();return function(){var i,o=t(n);if(r){var c=t(this).constructor;i=Reflect.construct(o,arguments,c)}else i=o.apply(this,arguments);return e(this,i)}}function rn(n,t,e){var r=n.slice();return r[3]=t[e].name,r[4]=t[e].content,r[6]=e,r}function on(n){var t,e,r;return{c:function(){t=s("meta"),this.h()},l:function(n){t=u(n,"META",{name:!0,content:!0}),this.h()},h:function(){g(t,"name",e=n[3]),g(t,"content",r=n[4])},m:function(n,e){p(n,t,e)},p:function(n,i){4&i&&e!==(e=n[3])&&g(t,"name",e),4&i&&r!==(r=n[4])&&g(t,"content",r)},d:function(n){n&&v(t)}}}function cn(n){var t,e,r,i,o;document.title=t=z.name+" • "+n[0];for(var c=n[2],a=[],s=0;s<c.length;s+=1)a[s]=on(rn(n,c,s));return i=new tn({props:{sections:n[1],project:"zitadel/site",dir:n[0]}}),{c:function(){for(var n=0;n<a.length;n+=1)a[n].c();e=C(),r=l(),y(i.$$.fragment)},l:function(n){for(var t=V('[data-svelte="svelte-4bltt5"]',document.head),o=0;o<a.length;o+=1)a[o].l(t);e=C(),t.forEach(v),r=h(n),E(i.$$.fragment,n)},m:function(n,t){for(var c=0;c<a.length;c+=1)a[c].m(document.head,null);$(document.head,e),p(n,r,t),_(i,n,t),o=!0},p:function(n,r){var s=A(r,1)[0];if((!o||1&s)&&t!==(t=z.name+" • "+n[0])&&(document.title=t),4&s){var l;for(c=n[2],l=0;l<c.length;l+=1){var u=rn(n,c,l);a[l]?a[l].p(u,s):(a[l]=on(u),a[l].c(),a[l].m(e.parentNode,e))}for(;l<a.length;l+=1)a[l].d(1);a.length=c.length}var f={};2&s&&(f.sections=n[1]),1&s&&(f.dir=n[0]),i.$set(f)},i:function(n){o||(w(i.$$.fragment,n),o=!0)},o:function(n){b(i.$$.fragment,n),o=!1},d:function(n){S(a,n),v(e),n&&v(r),R(i,n)}}}function an(n){return sn.apply(this,arguments)}function sn(){return(sn=k(P.mark((function n(t){var e,r,i,o;return P.wrap((function(n){for(;;)switch(n.prev=n.next){case 0:return e=t.params,e.lang,r=e.slug,n.next=4,this.fetch("".concat(r,".json")).then((function(n){return n.json()}));case 4:return i=n.sent,o=[],n.abrupt("return",{sections:i,slug:r,tags:o});case 7:case"end":return n.stop()}}),n,this)})))).apply(this,arguments)}function ln(n,t,e){var r=t.slug,i=t.sections,o=t.tags;return n.$set=function(n){"slug"in n&&e(0,r=n.slug),"sections"in n&&e(1,i=n.sections),"tags"in n&&e(2,o=n.tags)},[r,i,o]}var un=function(t){n(s,a);var e=en(s);function s(n){var t;return r(this,s),t=e.call(this),i(c(t),n,ln,cn,o,{slug:0,sections:1,tags:2}),t}return s}();export default un;export{an as preload};
