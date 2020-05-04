FROM ubuntu:latest

ENV ARCH amd64
ENV GOVERSION 1.7.5
ENV HOME /root
ENV PATH $PATH:/usr/local/go/bin
ENV LANG C.UTF-8
ENV LANGUAGE en_US:
ENV GOPATH $HOME
RUN apt-get update -y \
    && apt-get install vim curl git tmux tree  -y \
    && apt-get install wget \
    # vimの設定
    && echo "set number" >> ~/.vimrc \ 
    && echo "set fenc=utf-8" >> ~/.vimrc \ 
    && echo "set autoread" >> ~/.vimrc \ 
    && echo "set expandtab" >> ~/.vimrc \ 
    && echo "set hlsearch" >> ~/.vimrc \ 
    && echo "set ignorecase" >> ~/.vimrc \ 
    && echo "set incsearch" >> ~/.vimrc \ 
    && echo "set laststatus=2" >> ~/.vimrc \ 
    && echo "syntax on" >> ~/.vimrc \ 
    && echo "set autoindent" >> ~/.vimrc \ 
    && wget https://dl.google.com/go/go1.14.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.14.linux-amd64.tar.gz \
    && export PATH=$PATH:/usr/local/go/bin \
    && rm ./go1.14.linux-amd64.tar.gz \
    && curl -s -fLo /root/.vim/autoload/plug.vim --create-dirs https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim \
    && echo "let g:go_fmt_command = \"goimports\"" >> $HOME/.vimrc \
    && echo "filetype plugin indent on" >> $HOME/.vimrc \
    && git clone https://github.com/fatih/vim-go.git $HOME/.vim/pack/plugins/start/vim-go 
# vimで開発するなら::GoInstallBinaries

WORKDIR $HOME/work
