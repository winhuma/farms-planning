import{a as u}from"./axios-4a70c6fc.js";import{_ as d,o as i,c as _,a as e,w as n,v as a}from"./index-0fac0d14.js";const c={data(){return{area_id:void 0,num_person:void 0,num_day:void 0,cal_result:""}},methods:{handler(){typeof this.area_id!="number"||typeof this.num_person!="number"||typeof this.num_day!="number"?this.cal_result="FUCK YOU BITCH !!! you not enter any input":u.post("http://vps1225.vpshispeed.net:8080/waters/calculate/area",{area_id:this.area_id,number_person:this.num_person,number_day:this.num_day}).then(l=>{console.log(l.data),this.cal_result=l.data.result})}}},m={class:"area"},p=e("h1",null,"Calculate Water Require",-1),f=e("br",null,null,-1),b={class:"row g-3"},h={class:"col-md-4"},y=e("label",{class:"form-label"},"Area ID",-1),v={class:"col-md-4"},x=e("label",{class:"form-label"},"Number of person",-1),w={class:"col-md-4"},V=e("label",{class:"form-label"},"Number of day",-1),C={class:"input-group mb-3"},U=["value"];function g(l,t,B,k,s,r){return i(),_("div",m,[p,f,e("form",b,[e("div",h,[y,n(e("input",{type:"number",class:"form-control","onUpdate:modelValue":t[0]||(t[0]=o=>s.area_id=o)},null,512),[[a,s.area_id]])]),e("div",v,[x,n(e("input",{type:"number",class:"form-control","onUpdate:modelValue":t[1]||(t[1]=o=>s.num_person=o)},null,512),[[a,s.num_person]])]),e("div",w,[V,n(e("input",{type:"number",class:"form-control","onUpdate:modelValue":t[2]||(t[2]=o=>s.num_day=o)},null,512),[[a,s.num_day]])]),e("div",C,[e("button",{onClick:t[3]||(t[3]=(...o)=>r.handler&&r.handler(...o)),class:"btn btn-success",type:"button",id:"bt-summit1"},"Calculate"),e("input",{id:"input-summit1",type:"text",value:s.cal_result,class:"form-control",placeholder:"result","aria-label":"Example text with button addon","aria-describedby":"button-addon1",disabled:""},null,8,U)])])])}const D=d(c,[["render",g]]);export{D as default};
