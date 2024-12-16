import{p as u,c as f,b as E,a as l,r as lt,d as q,e as ut,t as Y,f as B,g as ft,s as dt,m as yt,u as mt}from"./props-CNBDvcf0.js";import{K as gt,L as pt,J as _t,a1 as bt,a2 as xt,g as T,a3 as D,h as I,k as S,i as m,j as N,p as w,u as F,f as d,d as O,T as K,t as vt,e as kt}from"./runtime-Dcs_0GP0.js";import{O as Ct}from"./index-D-8MO0q_.js";import{y as Et,g as M}from"./index-DPS9-N-h.js";let St=!1;function Nt(e,t,...r){var s=e,n=bt,i;gt(()=>{n!==(n=t())&&(i&&(xt(i),i=null),i=_t(()=>n(s,...r)))},pt)}function ht(e,t){var r=e.__className,s=Tt(t);(r!==s||St)&&(t==null?e.removeAttribute("class"):e.className=s,e.__className=s)}function Tt(e){return e??""}function P(e,t,r){if(r){if(e.classList.contains(t))return;e.classList.add(t)}else{if(!e.classList.contains(t))return;e.classList.remove(t)}}function wt(e,t,r,s){var n=e.__styles??(e.__styles={});n[t]!==r&&(n[t]=r,r==null?e.style.removeProperty(t):e.style.setProperty(t,r,""))}const Ot=e=>e.split("-").map(t=>t.charAt(0).toUpperCase()+t.slice(1)).join(""),At=e=>Ct(e.replace(/([A-Z])/g," $1").trim()),z=e=>Ot(At(e)),h="storybook-stories-extractor-context";function U(e){let t=u(e.isExtracting),r=u(e.register);return{get isExtracting(){return t},get register(){return r}}}function Rt(e){const{stories:t}=e,r=U({isExtracting:!0,register:s=>{t.set(s.exportName??z(s.name),s)}});T(h,r)}function Ft(){return D(h)||T(h,U({isExtracting:!1,register:()=>{}})),I(h)}const L="storybook-story-renderer-context";function Lt(e){let t=S(u(e.currentStoryExportName)),r=S(u(e.args)),s=S(u(e.storyContext));function n(i){N(t,u(i.currentStoryExportName)),N(r,u(i.args)),N(s,u(i.storyContext))}return{get args(){return m(r)},get storyContext(){return m(s)},get currentStoryExportName(){return m(t)},set:n}}function qt(){const e=Lt({currentStoryExportName:void 0,args:{},storyContext:{}});T(L,e)}function W(){return D(L)||qt(),I(L)}const R="storybook-stories-template-snippet-context";function Dt(){let e=S(void 0);function t(r){N(e,u(r))}return{get template(){return m(e)},set:t}}function It(){return D(R)||T(R,Dt()),I(R).template}var jt=Y('<p>No story rendered. See <a href="https://github.com/storybookjs/addon-svelte-csf#defining-stories" target="_blank">the docs</a> on how to define stories.</p>');function X(e,t){w(t,!0);const r=lt(t,["$$slots","$$events","$$legacy","children","name","exportName","play"]),s=t.exportName??z(t.name),n=Ft(),i=W(),a=It(),o=K(()=>!n.isExtracting&&i.currentStoryExportName===s);n.isExtracting&&n.register({children:t.children,name:t.name,exportName:s,play:t.play,...r});function c(g,x){x&&g.playFunction&&(g.playFunction.__play=x)}F(()=>{m(o)&&c(i.storyContext,t.play)});var y=f(),k=d(y);{var Q=g=>{var x=f(),V=d(x);{var $=p=>{var v=f(),A=d(v);Nt(A,()=>t.children,()=>i.args,()=>i.storyContext),l(p,v)},tt=p=>{var v=f(),A=d(v);{var et=_=>{a(_,()=>i.args,()=>i.storyContext)},rt=_=>{var j=f(),nt=d(j);{var st=b=>{var C=f(),ot=d(C);q(ot,()=>i.storyContext.component,(at,ct)=>{ct(at,ut(()=>i.args))}),l(b,C)},it=b=>{var C=jt();l(b,C)};E(nt,b=>{i.storyContext.component?b(st):b(it,!1)},!0)}l(_,j)};E(A,_=>{a?_(et):_(rt,!1)},!0)}l(p,v)};E(V,p=>{t.children?p($):p(tt,!1)})}l(g,x)};E(k,g=>{m(o)&&g(Q)})}l(e,y),O()}X.__docgen={data:[],name:"Story.svelte"};function ne(e){return{Story:X,meta:e}}var Bt=Y('<button type="button"> </button>');function Mt(e,t){w(t,!0);const r=B(t,"primary",3,!1),s=B(t,"size",3,"medium");var n=Bt();const i=K(()=>["storybook-button",`storybook-button--${s()}`].join(" "));n.__click=function(...o){var c;(c=t.onClick)==null||c.apply(this,o)};var a=kt(n);vt(()=>{ht(n,m(i)),P(n,"storybook-button--primary",r()),P(n,"storybook-button--secondary",!r()),wt(n,"background-color",t.backgroundColor),dt(a,t.label)}),l(e,n),O()}ft(["click"]);Mt.__docgen={data:[{name:"primary",visibility:"public",description:"Is this the principal call to action on the page?",keywords:[],kind:"let",type:{kind:"type",type:"boolean",text:"boolean"},static:!1,readonly:!1,defaultValue:"false"},{name:"backgroundColor",visibility:"public",description:"What background color to use",keywords:[],kind:"let",type:{kind:"type",type:"string",text:"string"},static:!1,readonly:!1},{name:"size",visibility:"public",description:"How large should the button be?",keywords:[],kind:"let",type:{kind:"union",type:[{kind:"const",type:"string",value:"small",text:'"small"'},{kind:"const",type:"string",value:"medium",text:'"medium"'},{kind:"const",type:"string",value:"large",text:'"large"'}],text:'"small" | "medium" | "large"'},static:!1,readonly:!1,defaultValue:'"medium"'},{name:"label",visibility:"public",description:"Button contents",keywords:[{name:"required",description:""}],kind:"let",type:{kind:"type",type:"string",text:"string"},static:!1,readonly:!1},{name:"onClick",visibility:"public",description:"The onclick event handler",keywords:[],kind:"let",type:{kind:"function",text:"() => void"},static:!1,readonly:!1}],name:"Button.svelte"};function G(e,t){w(t,!0),Rt(t.repository());var r=f(),s=d(r);q(s,()=>t.Stories,(n,i)=>{i(n,{})}),l(e,r),O()}G.__docgen={data:[{name:"Stories",visibility:"public",keywords:[{name:"required",description:""}],kind:"let",type:{kind:"function",text:"Component<{}, {}, string>"},static:!1,readonly:!1},{name:"repository",visibility:"public",keywords:[{name:"required",description:""}],kind:"let",type:{kind:"function",text:"() => StoriesRepository<Cmp>"},static:!1,readonly:!1}],name:"StoriesExtractor.svelte"};function Pt(e){switch(typeof e){case"number":case"symbol":return!1;case"string":return e.includes(".")||e.includes("[")||e.includes("]")}}function Yt(e){return Object.is(e,-0)?"-0":e.toString()}function Kt(e){const t=[],r=e.length;if(r===0)return t;let s=0,n="",i="",a=!1;for(e.charCodeAt(0)===46&&(t.push(""),s++);s<r;){const o=e[s];i?o==="\\"&&s+1<r?(s++,n+=e[s]):o===i?i="":n+=o:a?o==='"'||o==="'"?i=o:o==="]"?(a=!1,t.push(n),n=""):n+=o:o==="["?(a=!0,n&&(t.push(n),n="")):o==="."?n&&(t.push(n),n=""):n+=o,s++}return n&&t.push(n),t}function J(e,t,r){if(e==null)return r;switch(typeof t){case"string":{const s=e[t];return s===void 0?Pt(t)?J(e,Kt(t),r):r:s}case"number":case"symbol":{typeof t=="number"&&(t=Yt(t));const s=e[t];return s===void 0?r:s}default:{if(Array.isArray(t))return zt(e,t,r);Object.is(t==null?void 0:t.valueOf(),-0)?t="-0":t=String(t);const s=e[t];return s===void 0?r:s}}}function zt(e,t,r){if(t.length===0)return r;let s=e;for(let n=0;n<t.length;n++){if(s==null)return r;s=s[t[n]]}return s===void 0?r:s}const{addons:Ut}=__STORYBOOK_MODULE_PREVIEW_API__,Wt=Ut.getChannel(),Xt=e=>{const{storyContext:t}=e;if(Gt(t))return;const r=Jt({code:t.parameters.__svelteCsf.rawCode,args:e.args});setTimeout(()=>{Wt.emit(Et,{id:t.id,args:t.unmappedArgs,source:r})})},Gt=e=>{var n;const t=(n=e==null?void 0:e.parameters.docs)==null?void 0:n.source,r=e==null?void 0:e.parameters.__isArgsStory;return(e==null?void 0:e.parameters.__svelteCsf.rawCode)?(t==null?void 0:t.type)===M.DYNAMIC?!1:!r||(t==null?void 0:t.code)||(t==null?void 0:t.type)===M.CODE:!0},Jt=({code:e,args:t})=>{const r=Object.entries(t??{}).map(([i,a])=>Zt(i,a)).filter(i=>i);let s=r.join(" ");return s.length>50&&(s=`
  ${r.join(`
  `)}
`),e.replaceAll("{...args}",s).replace(/args(?:[\w\d_$\.\?\[\]"'])+/g,i=>{const a=i.replaceAll("?",""),o=J({args:t},a);return H(o)})},Ht=e=>{var r;const t=((r=e.getMockName)==null?void 0:r.call(e))??e.name;return t&&t!=="spy"?t:"() => {}"},H=e=>{var t;return typeof e=="object"&&e[Symbol.for("svelte.snippet")]?"snippet":typeof e=="function"?Ht(e):(t=JSON.stringify(e,null,1))==null?void 0:t.replace(/\n/g,"").replace(new RegExp("(?<!\\s)([}\\]])$")," $1")},Zt=(e,t)=>{if(t==null)return null;if(t===!0)return e;const r=H(t);return typeof t=="string"?`${e}=${r}`:`${e}={${r}}`};function Z(e,t){w(t,!0);const r=W();F(()=>{r.set({currentStoryExportName:t.exportName,args:t.args,storyContext:t.storyContext})}),F(()=>{Xt({args:t.args,storyContext:t.storyContext})});var s=f(),n=d(s);q(n,()=>t.Stories,(i,a)=>{a(i,{})}),l(e,s),O()}Z.__docgen={data:[{name:"Stories",visibility:"public",keywords:[{name:"required",description:""}],kind:"let",type:{kind:"function",text:"Component<{}, {}, string>"},static:!1,readonly:!1},{name:"exportName",visibility:"public",keywords:[{name:"required",description:""}],kind:"let",type:{kind:"type",type:"string",text:"string"},static:!1,readonly:!1},{name:"args",visibility:"public",keywords:[{name:"required",description:""}],kind:"let",type:{kind:"type",type:"any",text:"any"},static:!1,readonly:!1},{name:"storyContext",visibility:"public",keywords:[{name:"required",description:""}],kind:"let",type:{kind:"type",type:"any",text:"any"},static:!1,readonly:!1}],name:"StoryRenderer.svelte"};const{logger:Qt}=__STORYBOOK_MODULE_CLIENT_LOGGER__,Vt=document.createDocumentFragment?()=>document.createDocumentFragment():()=>document.createElement("div"),se=(e,t)=>{const r={stories:new Map};try{const n=yt(G,{target:Vt(),props:{Stories:e,repository:()=>r}});mt(n)}catch(n){Qt.error(`Error in mounting stories ${n.toString()}`,n)}const s={};for(const[n,i]of r.stories){const a={...i,render:(c,y)=>({Component:Z,props:{exportName:n,Stories:e,storyContext:y,args:c}})},o=t.play??i.play;o&&(a.play=c=>{var k;const y=(k=c.playFunction)==null?void 0:k.__play;return y?y(c):o(c)}),s[n]=a}return s};export{Mt as B,se as c,ne as d};