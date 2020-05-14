function GoRunWeb()
  copen
  let startCommand = "./webstart.sh"
  execute "AsyncRun " . startCommand
endfunction

function GoStopWeb()
  execute "AsyncStop"
endfunction

function GoRebuildWeb()
  execute "!" . "./webrebuild.sh"
endfunction

function GoDebugWeb()
  execute "!" . "./webdebug.sh"
endfunction

nnoremap <leader>pr :call GoRunWeb()<CR>
nnoremap <leader>ps :call GoStopWeb()<CR>
nnoremap <leader>pb :call GoRebuildWeb()<CR>
nnoremap <leader>pd :call GoDebugWeb()<CR>
