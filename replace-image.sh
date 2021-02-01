docker rm -f wiktionary-combine;
docker rmi alaverydev/audio-language-wiktionary-combine;
docker build -t alaverydev/audio-language-wiktionary-combine .
docker push alaverydev/audio-language-wiktionary-combine