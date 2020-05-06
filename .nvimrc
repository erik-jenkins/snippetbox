function GoRunWeb()
  copen
  let startCommand = "./startdev.sh"
  execute "AsyncRun " . startCommand
endfunction

function GoStopWeb()
  execute "AsyncStop"
endfunction

nnoremap <leader>pr :call GoRunWeb()<CR>
nnoremap <leader>ps :call GoStopWeb()<CR>
