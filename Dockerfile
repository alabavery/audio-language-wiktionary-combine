FROM golang:alpine3.13

RUN  mkdir wds_dir_mt && mkdir def_dir_mt && mkdir lemma_dir_mt && mkdir tgt_dir_mt
COPY ./src /src
WORKDIR /src
ENTRYPOINT go run main.go -words='/wds_dir_mt/words.json' -definitions=/def_dir_mt -lemmas=/lemma_dir_mt -target=/tgt_dir_mt
