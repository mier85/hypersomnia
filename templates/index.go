// Code generated by go generate; DO NOT EDIT.
package templates

const Index = `<!doctype html><html lang=en><meta charset=utf-8><meta name=viewport content="width=device-width,initial-scale=1,shrink-to-fit=no"><link rel=stylesheet href=https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css><link rel=stylesheet href=https://cdnjs.cloudflare.com/ajax/libs/jsoneditor/6.1.0/jsoneditor.min.css><script src=https://kit.fontawesome.com/43cf7afab1.js></script><title>Hypersomnia</title><style>div.active{background-color:#444;color:#fff;-webkit-border-radius:3px;-moz-border-radius:3px;border-radius:3px}.ace-jsoneditor.ace_editor{font-size:10px!important}#request-editor,#response-editor{height:100%}</style><div class="container-fluid p-2 pl-4">Environment: <select class=js-environment>
{{range .Envs}}
<option value={{.}}>{{.}}
{{end}}</select><div class=row><div class=col-sm><div class="mt-1 mb-1">Services:</div><div class=js-services></div></div><div class=col-sm><div class="mb-2 clearfix"><div class="js-active-endpoint float-left pt-1 pb-1"></div><button class="btn btn-sm btn-primary js-send float-right">Send</button>
<button class="btn btn-sm btn-secondary js-reset float-right mr-2">Reset</button></div><div id=request-editor></div></div><div class=col-sm><div class="mb-2 mt-3 clearfix"><span class="badge badge-secondary js-response-time float-right">...</span>
<span class="badge badge-secondary js-response-took float-right mr-2">...</span></div><div id=response-editor></div></div></div></div><script src=https://code.jquery.com/jquery-3.3.1.min.js crossorigin=anonymous></script><script src=https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js></script><script src=https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js></script><script src=https://momentjs.com/downloads/moment.min.js></script><script src=https://cdnjs.cloudflare.com/ajax/libs/jsoneditor/6.1.0/jsoneditor.min.js></script><script src=https://unpkg.com/dexie@2.0.4/dist/dexie.js></script><script src=https://cdn.jsdelivr.net/npm/jsonpath@1.0.2/jsonpath.min.js></script><script>String.prototype.replaceAll=function(search,replacement){let target=this;return target.replace(new RegExp(search,'g'),replacement);};function pregQuote(str,delimiter){return(str+'').replace(new RegExp('[.\\\\+*?\\[\\^\\]$(){}=!<>|:\\'+(delimiter||'/')+'-]','g'),'\\$&');}
function getMatches(string){let matches=[];let match;let regex=/Response\.(.+?)\((.+?),?(\s?int)?\)/g;while(match=regex.exec(string)){matches[match[0]]={endpoint:match[1],path:match[2],type:match[3]||'string',};}
return matches;}
function pack(v){let m={};for(let i in v.values){if(v.values[i].values==null||v.values[i].values.length==0){if(v.values[i].type=='string'){m[v.values[i].name]="";}else{m[v.values[i].name]=0;}}else{m[v.values[i].name]=pack(v.values[i]);}}
return m}
let storage=window.localStorage;let db=new Dexie("hypersomnia");db.version(1).stores({requests:'endpoint,body',responses:'endpoint,time,receivedAt,body'});let activeEndpoint=storage.getItem('active-endpoint');$(function(){let requestEditor=new JSONEditor(document.getElementById("request-editor"),{enableSort:false,enableTransform:false,mode:'code',});let responseEditor=new JSONEditor(document.getElementById("response-editor"),{mode:'code',});$('.js-environment').on('change',function(){$('.js-environment').prop('disabled',true)
$.ajax({method:'POST',url:'services',dataType:'json',contentType:'application/json',data:JSON.stringify({environment:$(this).val(),}),success:function(response){$('.js-services').text('');for(let i in response){let name=response[i].name;let id=name.replace(/\./g,'_');$('.js-services').append('<div style="cursor:pointer;" class="mt-1 mb-1 collapsed" data-toggle="collapse" href="#'+id+'"\n'+
'role="button">\n'+
'<i class="fas fa-cube pr-1"></i>\n'+
'<strong>'+name+'</strong>\n'+
'</div>\n'+
'<div class="mb-4 collapse js-endpoints" id="'+id+'" data-service="'+name+'">...</div>');}
$('.collapse').each(function(){let id=$(this).attr('id');if(storage.getItem(id+':show')==='true'){$(this).collapse('show');}else{$(this).collapse('hide');}});},error:function(){alert("AJAX error");},complete:function(){$('.js-environment').prop('disabled',false)}});});$('.js-environment').trigger('change');$(document).on('click','.js-endpoint-toggle',function(){let service=$(this).data('service');storage.setItem('active-service',service);let endpoint=$(this).data('endpoint');storage.setItem('active-endpoint',endpoint);$('.js-active-endpoint').text(endpoint);$('.js-endpoint-toggle').removeClass('active');$(this).addClass('active');let requestBody=JSON.parse($(this).find('.js-endpoint-request-template').text());requestEditor.set('loading...');db.requests.where('endpoint').equals(endpoint).first().then(function(request){if(request){requestBody=request.body;}}).catch(function(error){console.log(error);}).finally(function(){requestEditor.set(requestBody);});responseEditor.set('loading...');$('.js-response-took').text('...');$('.js-response-time').text('...');let cachedResponse={time:'...',receivedAt:'',body:''};db.responses.where('endpoint').equals(endpoint).first().then(function(response){if(response){cachedResponse=response;}}).catch(function(error){console.log(error);}).finally(function(){responseEditor.set(cachedResponse.body);$('.js-response-took').text(cachedResponse.time);$('.js-response-time').text(moment(cachedResponse.receivedAt).fromNow());});});$(document).on('keyup','#request-editor',function(){let endpoint=storage.getItem('active-endpoint');db.requests.put({endpoint:endpoint,body:requestEditor.get()});});$(document).on('click','.js-reset',function(){let endpoint=storage.getItem('active-endpoint');let requestTemplate=$('.js-endpoint-toggle[data-endpoint="'+endpoint+'"]').find('.js-endpoint-request-template').text();requestEditor.set(JSON.parse(requestTemplate));});$(document).on('show.bs.collapse','.collapse',function(){let id=$(this).attr('id');storage.setItem(id+':show',true);let serviceName=$(this).data('service');let container=$('.js-endpoints[id="'+id+'"]');container.html('loading...');$.ajax({method:'POST',url:'service',dataType:'json',contentType:'application/json',data:JSON.stringify({environment:$('.js-environment').val(),name:serviceName,}),success:function(response){container.html('<ul class="list-unstyled"></ul>');for(let i in response.endpoints){container.find('ul').append('<li><div class="ml-3 pl-1 pr-1 mb-2 js-endpoint-toggle"\n'+
'style="cursor: pointer;display: inline-block;"\n'+
'data-service="'+serviceName+'" data-endpoint="'+response.endpoints[i].name+'">\n'+
response.endpoints[i].name+'\n'+
'<pre style="display:none;"\n'+
'class="js-endpoint-request-template">'+JSON.stringify(pack(response.endpoints[i].request))+'</pre>\n'+
'</div></li>');}
if(activeEndpoint){console.log(activeEndpoint);$('.js-endpoint-toggle[data-endpoint="'+activeEndpoint+'"]').trigger('click');}},error:function(){alert("AJAX error");},complete:function(){}});}).on('hide.bs.collapse','.collapse',function(){let id=$(this).attr('id');storage.setItem(id+':show',false);});$('.collapse').each(function(){let id=$(this).attr('id');if(storage.getItem(id+':show')==='true'){$(this).collapse('show');}else{$(this).collapse('hide');}});$(document).on('click','.js-send',function(){$('.js-send').prop('disabled',true);let environment=$('.js-environment').val();let service=storage.getItem('active-service');let endpoint=storage.getItem('active-endpoint');let body;try{body=requestEditor.get();}catch(e){alert('Request: '+e.message);}
responseEditor.set('loading...')
$('.js-response-time').text('...');$('.js-response-took').text('...');let bodyText=JSON.stringify(body);let matches=getMatches(bodyText);let promises=[];for(let match in matches){let source=matches[match];let replace='';promises.push(db.responses.where('endpoint').equals(source.endpoint).first().then(function(response){if(response){replace=jsonpath.value(response.body,source.path);}}).catch(function(error){console.log(error);}).finally(function(){if(source.type==='int'){match='"'+match+'"';}
bodyText=bodyText.replaceAll(pregQuote(match),replace);}));}
Promise.all(promises).then(function(){body=JSON.parse(bodyText);console.log('Environment: '+environment+'\nService: '+service+'\nEndpoint: '+endpoint+'\nRequest: '+JSON.stringify(body));$.ajax({method:'POST',url:'call',dataType:'json',contentType:'application/json',data:JSON.stringify({environment:environment,service:service,endpoint:endpoint,body:body,}),success:function(response){let body;try{body=JSON.parse(response.Body);}catch(e){body=response.Body;}
$('.js-response-time').text('just now');$('.js-response-took').text(response.Time);responseEditor.set(body);db.responses.put({endpoint:endpoint,receivedAt:moment().format(),time:response.Time,body:body});},error:function(){alert("AJAX error");},complete:function(){$('.js-send').prop('disabled',false);}});});});$('#request-editor').keydown(function(e){if(e.ctrlKey&&e.keyCode===13){$('.js-send').trigger('click');}});if(activeEndpoint){$('.js-endpoint-toggle[data-endpoint="'+activeEndpoint+'"]').trigger('click');}});</script>`
