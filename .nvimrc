function GoRunWeb()
  copen
  let startCommand = "./scripts/webstart.sh"
  execute "AsyncRun " . startCommand
endfunction

function GoStopWeb()
  execute "AsyncStop"
endfunction

function GoRebuildWeb()
  execute "!" . "./scripts/webrebuild.sh"
endfunction

function GoDebugWeb()
  execute "!" . "./scripts/webdebug.sh"
endfunction

nnoremap <leader>pr :call GoRunWeb()<CR>
nnoremap <leader>ps :call GoStopWeb()<CR>
nnoremap <leader>pb :call GoRebuildWeb()<CR>
nnoremap <leader>pd :call GoDebugWeb()<CR>
