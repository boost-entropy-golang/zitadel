import{S as e,i as t,s,e as l,B as n,a,b as r,d as i,C as o,f as c,w as h,g as u,n as f,z as d,j as g,l as m,D as p,t as v,k as $,m as w,o as E,h as b,p as z,E as y,F as x,A as I,G as k,H as _,I as N,J as j,K as A,r as D,v as C,y as S,L,M as T,N as B,O as P,P as V,x as U,Q as O,q as R}from"./client.72e3f8d7.js";function q(e){let t,s,g;return{c(){t=l("i"),s=n("use"),this.h()},l(e){t=a(e,"I",{class:!0,style:!0});var l=r(t);s=a(l,"use",{"xlink:href":!0},1),r(s).forEach(i),l.forEach(i),this.h()},h(){o(s,"xlink:href",g="#"+e[0]),c(t,"class",e[0]),h(t,"font-size",e[1]),h(t,"height",e[1])},m(e,l){u(e,t,l),f(t,s)},p(e,[l]){1&l&&g!==(g="#"+e[0])&&o(s,"xlink:href",g),1&l&&c(t,"class",e[0]),2&l&&h(t,"font-size",e[1]),2&l&&h(t,"height",e[1])},i:d,o:d,d(e){e&&i(t)}}}function F(e,t,s){let{name:l}=t,{size:n="20px"}=t;return e.$$set=e=>{"name"in e&&s(0,l=e.name),"size"in e&&s(1,n=e.size)},[l,n]}class M extends e{constructor(e){super(),t(this,e,F,q,s,{name:0,size:1})}}var H="ZITADEL Documentation";function K(e,t,s){const l=e.slice();return l[9]=t[s],l}function G(e,t,s){const l=e.slice();return l[12]=t[s],l}function Y(e){let t,s,n;return s=new M({props:{name:"las la-arrow-right"}}),{c(){t=l("div"),$(s.$$.fragment),this.h()},l(e){t=a(e,"DIV",{class:!0});var l=r(t);w(s.$$.fragment,l),l.forEach(i),this.h()},h(){c(t,"class","icon-container svelte-7kfwfv")},m(e,l){u(e,t,l),E(s,t,null),n=!0},i(e){n||(v(s.$$.fragment,e),n=!0)},o(e){b(s.$$.fragment,e),n=!1},d(e){e&&i(t),z(s)}}}function J(e){let t,s,n;return s=new M({props:{name:"las la-arrow-right"}}),{c(){t=l("div"),$(s.$$.fragment),this.h()},l(e){t=a(e,"DIV",{class:!0});var l=r(t);w(s.$$.fragment,l),l.forEach(i),this.h()},h(){c(t,"class","icon-container svelte-7kfwfv")},m(e,l){u(e,t,l),E(s,t,null),n=!0},i(e){n||(v(s.$$.fragment,e),n=!0)},o(e){b(s.$$.fragment,e),n=!1},d(e){e&&i(t),z(s)}}}function Q(e){let t,s,n,o,h,d,$=e[12].title+"",w=e[12].slug===e[3]&&J();return{c(){t=l("a"),n=g(),w&&w.c(),this.h()},l(e){t=a(e,"A",{class:!0,href:!0,"data-level":!0});var s=r(t);n=m(s),w&&w.l(s),s.forEach(i),this.h()},h(){s=new j(n),c(t,"class","subsection svelte-7kfwfv"),c(t,"href",o=e[1]+"#"+e[12].slug),c(t,"data-level",h=e[12].level),p(t,"active",e[12].slug===e[3])},m(e,l){u(e,t,l),s.m($,t),f(t,n),w&&w.m(t,null),d=!0},p(e,l){(!d||4&l)&&$!==($=e[12].title+"")&&s.p($),e[12].slug===e[3]?w?12&l&&v(w,1):(w=J(),w.c(),v(w,1),w.m(t,null)):w&&(y(),b(w,1,1,(()=>{w=null})),x()),(!d||6&l&&o!==(o=e[1]+"#"+e[12].slug))&&c(t,"href",o),(!d||4&l&&h!==(h=e[12].level))&&c(t,"data-level",h),12&l&&p(t,"active",e[12].slug===e[3])},i(e){d||(v(w),d=!0)},o(e){b(w),d=!1},d(e){e&&i(t),w&&w.d()}}}function W(e){let t,s,n,o,h,d,$,w,E=e[9].metadata.title+"",z=e[9].slug===e[3]&&Y(),k=e[9].subsections,_=[];for(let t=0;t<k.length;t+=1)_[t]=Q(G(e,k,t));const N=e=>b(_[e],1,1,(()=>{_[e]=null}));return{c(){t=l("li"),s=l("a"),o=g(),z&&z.c(),d=g();for(let e=0;e<_.length;e+=1)_[e].c();$=g(),this.h()},l(e){t=a(e,"LI",{class:!0});var l=r(t);s=a(l,"A",{class:!0,href:!0});var n=r(s);o=m(n),z&&z.l(n),n.forEach(i),d=m(l);for(let e=0;e<_.length;e+=1)_[e].l(l);$=m(l),l.forEach(i),this.h()},h(){n=new j(o),c(s,"class","section svelte-7kfwfv"),c(s,"href",h=e[1]+"#"+e[9].slug),p(s,"active",e[9].slug===e[3]),c(t,"class","svelte-7kfwfv")},m(e,l){u(e,t,l),f(t,s),n.m(E,s),f(s,o),z&&z.m(s,null),f(t,d);for(let e=0;e<_.length;e+=1)_[e].m(t,null);f(t,$),w=!0},p(e,l){if((!w||4&l)&&E!==(E=e[9].metadata.title+"")&&n.p(E),e[9].slug===e[3]?z?12&l&&v(z,1):(z=Y(),z.c(),v(z,1),z.m(s,null)):z&&(y(),b(z,1,1,(()=>{z=null})),x()),(!w||6&l&&h!==(h=e[1]+"#"+e[9].slug))&&c(s,"href",h),12&l&&p(s,"active",e[9].slug===e[3]),14&l){let s;for(k=e[9].subsections,s=0;s<k.length;s+=1){const n=G(e,k,s);_[s]?(_[s].p(n,l),v(_[s],1)):(_[s]=Q(n),_[s].c(),v(_[s],1),_[s].m(t,$))}for(y(),s=k.length;s<_.length;s+=1)N(s);x()}},i(e){if(!w){v(z);for(let e=0;e<k.length;e+=1)v(_[e]);w=!0}},o(e){b(z),_=_.filter(Boolean);for(let e=0;e<_.length;e+=1)b(_[e]);w=!1},d(e){e&&i(t),z&&z.d(),I(_,e)}}}function Z(e){let t,s,n,o,h=e[2],f=[];for(let t=0;t<h.length;t+=1)f[t]=W(K(e,h,t));const d=e=>b(f[e],1,1,(()=>{f[e]=null}));return{c(){t=l("ul");for(let e=0;e<f.length;e+=1)f[e].c();this.h()},l(e){t=a(e,"UL",{class:!0});var s=r(t);for(let e=0;e<f.length;e+=1)f[e].l(s);s.forEach(i),this.h()},h(){c(t,"class","reference-toc svelte-7kfwfv")},m(l,a){u(l,t,a);for(let e=0;e<f.length;e+=1)f[e].m(t,null);e[6](t),s=!0,n||(o=[k(t,"mouseenter",e[7]),k(t,"mouseleave",e[8])],n=!0)},p(e,[s]){if(14&s){let l;for(h=e[2],l=0;l<h.length;l+=1){const n=K(e,h,l);f[l]?(f[l].p(n,s),v(f[l],1)):(f[l]=W(n),f[l].c(),v(f[l],1),f[l].m(t,null))}for(y(),l=h.length;l<f.length;l+=1)d(l);x()}},i(e){if(!s){for(let e=0;e<h.length;e+=1)v(f[e]);s=!0}},o(e){f=f.filter(Boolean);for(let e=0;e<f.length;e+=1)b(f[e]);s=!1},d(s){s&&i(t),I(f,s),e[6](null),n=!1,_(o)}}}function X(e,t,s){let l,{dir:n=""}=t,{sections:a=[]}=t,{active_section:r=null}=t,{show_contents:i}=t,{prevent_sidebar_scroll:o=!1}=t;N((()=>{if(o||i&&window.innerWidth<832)return;const e=l.querySelector(".active");if(e){const{top:t,bottom:s}=e.getBoundingClientRect(),n=200,a=window.innerHeight-200;t>a?l.parentNode.scrollBy({top:t-a,left:0,behavior:"smooth"}):s<n&&l.parentNode.scrollBy({top:s-n,left:0,behavior:"smooth"})}}));return e.$$set=e=>{"dir"in e&&s(1,n=e.dir),"sections"in e&&s(2,a=e.sections),"active_section"in e&&s(3,r=e.active_section),"show_contents"in e&&s(5,i=e.show_contents),"prevent_sidebar_scroll"in e&&s(0,o=e.prevent_sidebar_scroll)},[o,n,a,r,l,i,function(e){A[e?"unshift":"push"]((()=>{l=e,s(4,l)}))},()=>s(0,o=!0),()=>s(0,o=!1)]}class ee extends e{constructor(e){super(),t(this,e,X,Z,s,{dir:1,sections:2,active_section:3,show_contents:5,prevent_sidebar_scroll:0})}}function te(e,t,s){const l=e.slice();return l[10]=t[s],l[12]=s,l}function se(e){let t,s,n=e[10].slug+"";return{c(){t=l("p"),s=D(n),this.h()},l(e){t=a(e,"P",{class:!0});var l=r(t);s=C(l,n),l.forEach(i),this.h()},h(){c(t,"class","title svelte-1rbeejz")},m(e,l){u(e,t,l),f(t,s)},p(e,t){4&t&&n!==(n=e[10].slug+"")&&S(s,n)},d(e){e&&i(t)}}}function le(e){let t,s,n,o,h=e[10].title+"",d=e[10].parent+"";return{c(){t=l("p"),s=D(h),n=l("span"),o=D(d),this.h()},l(e){t=a(e,"P",{class:!0});var l=r(t);s=C(l,h),n=a(l,"SPAN",{class:!0});var c=r(n);o=C(c,d),c.forEach(i),l.forEach(i),this.h()},h(){c(n,"class","second-param svelte-1rbeejz"),c(t,"class","title svelte-1rbeejz")},m(e,l){u(e,t,l),f(t,s),f(t,n),f(n,o)},p(e,t){4&t&&h!==(h=e[10].title+"")&&S(s,h),4&t&&d!==(d=e[10].parent+"")&&S(o,d)},d(e){e&&i(t)}}}function ne(e){let t,s,n,o,d,p,v,$,w,E,b,z,y,x,I,_,N=e[10].slug+"",j=(e[10].html||e[10].slug)+"";function A(e,t){return e[10].level>2?le:se}let L=A(e),T=L(e);return{c(){t=l("a"),s=l("div"),T.c(),n=g(),o=l("p"),d=D(e[0]),p=D("#"),v=D(N),$=g(),w=l("p"),E=D(j),b=g(),z=l("i"),y=g(),this.h()},l(l){t=a(l,"A",{tabindex:!0,class:!0,href:!0,id:!0});var c=r(t);s=a(c,"DIV",{class:!0});var h=r(s);T.l(h),n=m(h),o=a(h,"P",{class:!0,style:!0});var u=r(o);d=C(u,e[0]),p=C(u,"#"),v=C(u,N),u.forEach(i),$=m(h),w=a(h,"P",{class:!0});var f=r(w);E=C(f,j),f.forEach(i),h.forEach(i),b=m(c),z=a(c,"I",{class:!0}),r(z).forEach(i),y=m(c),c.forEach(i),this.h()},h(){c(o,"class","desc svelte-1rbeejz"),h(o,"color","#85d996"),c(w,"class","desc svelte-1rbeejz"),c(s,"class","text svelte-1rbeejz"),c(z,"class","las la-link svelte-1rbeejz"),c(t,"tabindex","0"),c(t,"class","result-item svelte-1rbeejz"),c(t,"href",x=e[0]+"#"+e[10].slug),c(t,"id",e[12])},m(l,a){u(l,t,a),f(t,s),T.m(s,null),f(s,n),f(s,o),f(o,d),f(o,p),f(o,v),f(s,$),f(s,w),f(w,E),f(t,b),f(t,z),f(t,y),I||(_=k(t,"click",e[4]),I=!0)},p(e,l){L===(L=A(e))&&T?T.p(e,l):(T.d(1),T=L(e),T&&(T.c(),T.m(s,n))),1&l&&S(d,e[0]),4&l&&N!==(N=e[10].slug+"")&&S(v,N),4&l&&j!==(j=(e[10].html||e[10].slug)+"")&&S(E,j),5&l&&x!==(x=e[0]+"#"+e[10].slug)&&c(t,"href",x)},d(e){e&&i(t),T.d(),I=!1,_()}}}function ae(e){let t,s,n,o,h,p,v,$,w,E,b,z,y,x,N=e[2],j=[];for(let t=0;t<N.length;t+=1)j[t]=ne(te(e,N,t));return{c(){t=l("div"),s=g(),n=l("div"),o=l("div"),h=l("i"),p=g(),v=l("input"),$=g(),w=l("p"),E=D("Search Results:"),b=g(),z=l("div");for(let e=0;e<j.length;e+=1)j[e].c();this.h()},l(e){t=a(e,"DIV",{class:!0}),r(t).forEach(i),s=m(e),n=a(e,"DIV",{class:!0});var l=r(n);o=a(l,"DIV",{class:!0});var c=r(o);h=a(c,"I",{class:!0}),r(h).forEach(i),p=m(c),v=a(c,"INPUT",{placeholder:!0,class:!0}),c.forEach(i),$=m(l),w=a(l,"P",{class:!0});var u=r(w);E=C(u,"Search Results:"),u.forEach(i),b=m(l),z=a(l,"DIV",{tabindex:!0,class:!0});var f=r(z);for(let e=0;e<j.length;e+=1)j[e].l(f);f.forEach(i),l.forEach(i),this.h()},h(){c(t,"class","overlay svelte-1rbeejz"),c(h,"class","las la-search svelte-1rbeejz"),c(v,"placeholder","Search for something"),c(v,"class","svelte-1rbeejz"),c(o,"class","search-line svelte-1rbeejz"),c(w,"class","result-d svelte-1rbeejz"),c(z,"tabindex","-1"),c(z,"class","result-list svelte-1rbeejz"),c(n,"class","search-field svelte-1rbeejz")},m(l,a){u(l,t,a),u(l,s,a),u(l,n,a),f(n,o),f(o,h),f(o,p),f(o,v),L(v,e[1]),f(n,$),f(n,w),f(w,E),f(n,b),f(n,z);for(let e=0;e<j.length;e+=1)j[e].m(z,null);y||(x=[k(window,"keydown",e[3]),k(t,"click",e[4]),k(v,"input",e[7]),T(re.call(null,v))],y=!0)},p(e,[t]){if(2&t&&v.value!==e[1]&&L(v,e[1]),21&t){let s;for(N=e[2],s=0;s<N.length;s+=1){const l=te(e,N,s);j[s]?j[s].p(l,t):(j[s]=ne(l),j[s].c(),j[s].m(z,null))}for(;s<j.length;s+=1)j[s].d(1);j.length=N.length}},i:d,o:d,d(e){e&&i(t),e&&i(s),e&&i(n),I(j,e),y=!1,_(x)}}}function re(e){e.focus()}function ie(e,t,s){let{sections:l}=t,{slug:n}=t,a=[],r=0;const i=B();let o="";return e.$$set=e=>{"sections"in e&&s(5,l=e.sections),"slug"in e&&s(0,n=e.slug)},e.$$.update=()=>{64&e.$$.dirty&&function(e){console.log(e);const t=document.getElementById(e);t&&(console.log("focus: "+t),t.focus())}(r),2&e.$$.dirty&&function(e){const t=e.toLowerCase(),n=l.filter((e=>{const s=e.slug.toLowerCase().includes(t),l=e.html.replace(/<[^>]*>?/gm,"").toLowerCase().includes(t);return s||l})).map((e=>({title:e.slug,slug:e.slug}))),r=l.map((e=>e.subsections.map((t=>({parent:e.slug,...t}))))).flat().filter((e=>{if(e.slug){const s=e.slug.toLowerCase().includes(t),l=e.title.toLowerCase().includes(t);return s||l}}));s(2,a=n.concat(r)),console.log(a)}(o)},[n,o,a,function(e){console.log(e),e&&(37==e.keyCode||38==e.keyCode?r>0&&(e.preventDefault(),s(6,r--,r)):39!=e.keyCode&&40!=e.keyCode||(e.preventDefault(),s(6,r++,r)))},function(){i("close",{closed:!0})},l,r,function(){o=this.value,s(1,o)}]}class oe extends e{constructor(e){super(),t(this,e,ie,ae,s,{sections:5,slug:0})}}function ce(e){let t,s,n,o,h,p,v,$,w,E,b,z,y,x=(e[0]||"Ctrl")+"";return{c(){t=l("button"),s=l("i"),n=g(),o=l("span"),h=D("Search this site"),p=g(),v=l("span"),$=g(),w=l("span"),E=D(x),b=D(" F"),this.h()},l(e){t=a(e,"BUTTON",{class:!0});var l=r(t);s=a(l,"I",{class:!0}),r(s).forEach(i),n=m(l),o=a(l,"SPAN",{class:!0});var c=r(o);h=C(c,"Search this site"),c.forEach(i),p=m(l),v=a(l,"SPAN",{class:!0}),r(v).forEach(i),$=m(l),w=a(l,"SPAN",{class:!0});var u=r(w);E=C(u,x),b=C(u," F"),u.forEach(i),l.forEach(i),this.h()},h(){c(s,"class","las la-search svelte-5en90e"),c(o,"class","svelte-5en90e"),c(v,"class","fill-space svelte-5en90e"),c(w,"class","strg svelte-5en90e"),c(t,"class","search-trigger svelte-5en90e")},m(l,a){u(l,t,a),f(t,s),f(t,n),f(t,o),f(o,h),f(t,p),f(t,v),f(t,$),f(t,w),f(w,E),f(w,b),z||(y=k(t,"click",e[1]),z=!0)},p(e,[t]){1&t&&x!==(x=(e[0]||"Ctrl")+"")&&S(E,x)},i:d,o:d,d(e){e&&i(t),z=!1,y()}}}function he(e,t,s){let l="";return P((()=>{s(0,l=navigator.platform.indexOf("Mac")>-1?"Cmd":"Ctrl")})),[l,function(t){V(e,t)}]}class ue extends e{constructor(e){super(),t(this,e,he,ce,s,{})}}const{window:fe}=O;function de(e,t,s){const l=e.slice();return l[17]=t[s],l}function ge(e){let t,s,n,o,h,d,p,y,x,I,k,_,N,A,D,C,S,L,T,B=e[17].metadata.title+"",P=e[17].html+"";return N=new M({props:{name:"las la-external-link-alt",size:"24px"}}),{c(){t=l("section"),s=l("h2"),n=l("span"),h=g(),d=l("a"),y=g(),I=g(),k=l("small"),_=l("a"),$(N.$$.fragment),D=g(),S=g(),this.h()},l(e){t=a(e,"SECTION",{"data-id":!0,class:!0});var l=r(t);s=a(l,"H2",{class:!0});var o=r(s);n=a(o,"SPAN",{class:!0,id:!0}),r(n).forEach(i),h=m(o),d=a(o,"A",{href:!0,class:!0,"aria-hidden":!0}),r(d).forEach(i),y=m(o),I=m(o),k=a(o,"SMALL",{class:!0});var c=r(k);_=a(c,"A",{href:!0,title:!0,class:!0});var u=r(_);w(N.$$.fragment,u),u.forEach(i),c.forEach(i),o.forEach(i),D=m(l),S=m(l),l.forEach(i),this.h()},h(){c(n,"class","offset-anchor"),c(n,"id",o=e[17].slug),c(d,"href",p=e[3]+"#"+e[17].slug),c(d,"class","anchor"),c(d,"aria-hidden",""),x=new j(I),c(_,"href",A="https://github.com/"+e[0]+"/"+e[2]+"/edit/main/site/"+e[1]+"/"+e[3]+"/"+e[17].file),c(_,"title",e[4]),c(_,"class","svelte-him5z6"),c(k,"class","svelte-him5z6"),c(s,"class","svelte-him5z6"),C=new j(S),c(t,"data-id",L=e[17].slug),c(t,"class","svelte-him5z6")},m(e,l){u(e,t,l),f(t,s),f(s,n),f(s,h),f(s,d),f(s,y),x.m(B,s),f(s,I),f(s,k),f(k,_),E(N,_,null),f(t,D),C.m(P,t),f(t,S),T=!0},p(e,s){(!T||32&s&&o!==(o=e[17].slug))&&c(n,"id",o),(!T||40&s&&p!==(p=e[3]+"#"+e[17].slug))&&c(d,"href",p),(!T||32&s)&&B!==(B=e[17].metadata.title+"")&&x.p(B),(!T||47&s&&A!==(A="https://github.com/"+e[0]+"/"+e[2]+"/edit/main/site/"+e[1]+"/"+e[3]+"/"+e[17].file))&&c(_,"href",A),(!T||16&s)&&c(_,"title",e[4]),(!T||32&s)&&P!==(P=e[17].html+"")&&C.p(P),(!T||32&s&&L!==(L=e[17].slug))&&c(t,"data-id",L)},i(e){T||(v(N.$$.fragment,e),T=!0)},o(e){b(N.$$.fragment,e),T=!1},d(e){e&&i(t),z(N)}}}function me(e){let t,s;return t=new oe({props:{sections:e[5],slug:e[3]}}),t.$on("close",e[11]),{c(){$(t.$$.fragment)},l(e){w(t.$$.fragment,e)},m(e,l){E(t,e,l),s=!0},p(e,s){const l={};32&s&&(l.sections=e[5]),8&s&&(l.slug=e[3]),t.$set(l)},i(e){s||(v(t.$$.fragment,e),s=!0)},o(e){b(t.$$.fragment,e),s=!1},d(e){z(t,e)}}}function pe(e){let t,s,n,o,h,d,N,j,A,D,C,S,L,T,B,P,V,O,R=e[5],q=[];for(let t=0;t<R.length;t+=1)q[t]=ge(de(e,R,t));const F=e=>b(q[e],1,1,(()=>{q[e]=null}));j=new ue({}),j.$on("click",e[11]),D=new ee({props:{dir:e[3],sections:e[5],active_section:e[7],show_contents:e[10]}}),L=new M({props:{name:e[10]?"las la-times":"las la-bars"}});let H=1==e[6]&&me(e);return{c(){t=l("div");for(let e=0;e<q.length;e+=1)q[e].c();s=g(),n=l("div"),h=g(),d=l("aside"),N=l("div"),$(j.$$.fragment),A=g(),$(D.$$.fragment),C=g(),S=l("button"),$(L.$$.fragment),T=g(),H&&H.c(),B=U(),this.h()},l(e){t=a(e,"DIV",{class:!0});var l=r(t);for(let e=0;e<q.length;e+=1)q[e].l(l);l.forEach(i),s=m(e),n=a(e,"DIV",{class:!0}),r(n).forEach(i),h=m(e),d=a(e,"ASIDE",{class:!0});var o=r(d);N=a(o,"DIV",{class:!0});var c=r(N);w(j.$$.fragment,c),A=m(c),w(D.$$.fragment,c),c.forEach(i),C=m(o),S=a(o,"BUTTON",{class:!0});var u=r(S);w(L.$$.fragment,u),u.forEach(i),o.forEach(i),T=m(e),H&&H.l(e),B=U(),this.h()},h(){c(t,"class","content listify svelte-him5z6"),c(n,"class",o="overlay "+(e[10]?"visible":"")+" svelte-him5z6"),c(N,"class","sidebar svelte-him5z6"),c(S,"class","svelte-him5z6"),c(d,"class","sidebar-container svelte-him5z6"),p(d,"open",e[10])},m(l,a){u(l,t,a);for(let e=0;e<q.length;e+=1)q[e].m(t,null);e[13](t),u(l,s,a),u(l,n,a),u(l,h,a),u(l,d,a),f(d,N),E(j,N,null),f(N,A),E(D,N,null),f(d,C),f(d,S),E(L,S,null),e[16](d),u(l,T,a),H&&H.m(l,a),u(l,B,a),P=!0,V||(O=[k(fe,"keydown",e[12]),k(N,"click",e[14]),k(S,"click",e[15])],V=!0)},p(e,[s]){if(63&s){let l;for(R=e[5],l=0;l<R.length;l+=1){const n=de(e,R,l);q[l]?(q[l].p(n,s),v(q[l],1)):(q[l]=ge(n),q[l].c(),v(q[l],1),q[l].m(t,null))}for(y(),l=R.length;l<q.length;l+=1)F(l);x()}(!P||1024&s&&o!==(o="overlay "+(e[10]?"visible":"")+" svelte-him5z6"))&&c(n,"class",o);const l={};8&s&&(l.dir=e[3]),32&s&&(l.sections=e[5]),128&s&&(l.active_section=e[7]),1024&s&&(l.show_contents=e[10]),D.$set(l);const a={};1024&s&&(a.name=e[10]?"las la-times":"las la-bars"),L.$set(a),1024&s&&p(d,"open",e[10]),1==e[6]?H?(H.p(e,s),64&s&&v(H,1)):(H=me(e),H.c(),v(H,1),H.m(B.parentNode,B)):H&&(y(),b(H,1,1,(()=>{H=null})),x())},i(e){if(!P){for(let e=0;e<R.length;e+=1)v(q[e]);v(j.$$.fragment,e),v(D.$$.fragment,e),v(L.$$.fragment,e),v(H),P=!0}},o(e){q=q.filter(Boolean);for(let e=0;e<q.length;e+=1)b(q[e]);b(j.$$.fragment,e),b(D.$$.fragment,e),b(L.$$.fragment,e),b(H),P=!1},d(l){l&&i(t),I(q,l),e[13](null),l&&i(s),l&&i(n),l&&i(h),l&&i(d),z(j),z(D),z(L),e[16](null),l&&i(T),H&&H.d(l),l&&i(B),V=!1,_(O)}}}function ve(e,t,s){let l,n,a,{owner:r="caos"}=t,{path:i="docs"}=t,{project:o="zitadel"}=t,{dir:c=""}=t,{edit_title:h="edit this section"}=t,{sections:u}=t,f=!1,d=!1;P((()=>{const e=n.querySelectorAll("[id]:not([data-scrollignore])");let t;const a=()=>{const{top:s}=n.getBoundingClientRect();t=[].map.call(e,(e=>e.getBoundingClientRect().top-s))};let r=window.location.hash.slice(1);const i=()=>{const n=-window.scrollY;let a=e.length;for(;a--;)if(t[a]+n<40){const t=e[a],{id:n}=t;return void(n!==r&&(s(7,l=n),r=n))}};window.addEventListener("scroll",i,!0),window.addEventListener("resize",a,!0);const o=[setTimeout(a,1e3),setTimeout(i,5e3)];return a(),i(),()=>{window.removeEventListener("scroll",i,!0),window.removeEventListener("resize",a,!0),o.forEach((e=>clearTimeout(e)))}}));return e.$$set=e=>{"owner"in e&&s(0,r=e.owner),"path"in e&&s(1,i=e.path),"project"in e&&s(2,o=e.project),"dir"in e&&s(3,c=e.dir),"edit_title"in e&&s(4,h=e.edit_title),"sections"in e&&s(5,u=e.sections)},[r,i,o,c,h,u,f,l,n,a,d,function(e){s(6,f=!e.detail.closed)},function(e){const t=navigator.platform.indexOf("Mac")>-1?e.metaKey:e.ctrlKey;(114==e.keyCode||t&&70==e.keyCode)&&(e.preventDefault(),s(6,f=!f))},function(e){A[e?"unshift":"push"]((()=>{n=e,s(8,n)}))},()=>s(10,d=!1),()=>s(10,d=!d),function(e){A[e?"unshift":"push"]((()=>{a=e,s(9,a)}))}]}class $e extends e{constructor(e){super(),t(this,e,ve,pe,s,{owner:0,path:1,project:2,dir:3,edit_title:4,sections:5})}}function we(e){for(var t=function e(t,s){return t&&(s(t)?t:e(t.parentNode,s))},s=function(e){(e=e||window.event).preventDefault?e.preventDefault():e.returnValue=!1;var s=e.target||e.srcElement,n=t(s,(function(e){return e.tagName&&"FIGURE"===e.tagName.toUpperCase()}));if(n){for(var a,r=n.parentNode,i=n.parentNode.childNodes,o=i.length,c=0,h=0;h<o;h++)if(1===i[h].nodeType){if(i[h]===n){a=c;break}c++}return a>=0&&l(a,r),!1}},l=function(e,t,s,l){var n,a,r=document.querySelectorAll(".pswp")[0];if(a=function(e){for(var t,s,l,n,a=e.childNodes,r=a.length,i=[],o=0;o<r;o++){if(1!==(t=a[o]).nodeType)continue;const e=(s=t.children[0]).getAttribute("data-size");l=e?e.split("x"):"1920x1080".split("x"),n={src:s.getAttribute("href"),w:parseInt(l[0],10),h:parseInt(l[1],10)},t.children.length>1&&(n.title=t.children[1].innerHTML),s.children.length>0&&(n.msrc=s.children[0].getAttribute("src")),n.el=t,i.push(n)}return i}(t),n={galleryUID:t.getAttribute("data-pswp-uid"),getThumbBoundsFn:function(e){var t=a[e].el.getElementsByTagName("img")[0],s=window.pageYOffset||document.documentElement.scrollTop,l=t.getBoundingClientRect();return{x:l.left,y:l.top+s,w:l.width}}},l)if(n.galleryPIDs){for(var i=0;i<a.length;i++)if(a[i].pid==e){n.index=i;break}}else n.index=parseInt(e,10)-1;else n.index=parseInt(e,10);isNaN(n.index)||(s&&(n.showAnimationDuration=0),new PhotoSwipe(r,PhotoSwipeUI_Default,a,n).init())},n=document.querySelectorAll(e),a=0,r=n.length;a<r;a++)n[a].setAttribute("data-pswp-uid",a+1),n[a].onclick=s;var i=function(){var e=window.location.hash.substring(1),t={};if(e.length<5)return t;for(var s=e.split("&"),l=0;l<s.length;l++)if(s[l]){var n=s[l].split("=");n.length<2||(t[n[0]]=n[1])}return t.gid&&(t.gid=parseInt(t.gid,10)),t}();i.pid&&i.gid&&l(i.pid,n[i.gid-1],!0,!0)}function Ee(e){let t,s;return{c(){s=U(),this.h()},l(e){s=U(),this.h()},h(){t=new j(s)},m(l,n){t.m(e[2],l,n),u(l,s,n)},p(e,s){4&s&&t.p(e[2])},d(e){e&&i(s),e&&t.d()}}}function be(e){let t,s,l,n,a;document.title=t=H+" • "+e[0];let r=e[2]&&Ee(e);return n=new $e({props:{sections:e[1],dir:e[0]}}),{c(){r&&r.c(),s=U(),l=g(),$(n.$$.fragment)},l(e){const t=R('[data-svelte="svelte-18bu559"]',document.head);r&&r.l(t),s=U(),t.forEach(i),l=m(e),w(n.$$.fragment,e)},m(e,t){r&&r.m(document.head,null),f(document.head,s),u(e,l,t),E(n,e,t),a=!0},p(e,[l]){(!a||1&l)&&t!==(t=H+" • "+e[0])&&(document.title=t),e[2]?r?r.p(e,l):(r=Ee(e),r.c(),r.m(s.parentNode,s)):r&&(r.d(1),r=null);const i={};2&l&&(i.sections=e[1]),1&l&&(i.dir=e[0]),n.$set(i)},i(e){a||(v(n.$$.fragment,e),a=!0)},o(e){b(n.$$.fragment,e),a=!1},d(e){r&&r.d(e),i(s),e&&i(l),z(n,e)}}}async function ze({params:e}){const{slug:t}=e,{sections:s,seo:l}=await this.fetch(`${t}.json`).then((e=>e.json()));return{sections:s,seo:l,slug:t}}function ye(e,t,s){let{slug:l}=t,{sections:n}=t,{seo:a}=t;return P((()=>{we(".zitadel-gallery")})),e.$$set=e=>{"slug"in e&&s(0,l=e.slug),"sections"in e&&s(1,n=e.sections),"seo"in e&&s(2,a=e.seo)},[l,n,a]}export default class extends e{constructor(e){super(),t(this,e,ye,be,s,{slug:0,sections:1,seo:2})}}export{ze as preload};
