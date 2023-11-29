kafka-topics --bootstrap-server broker:9091 \
             --create \
             --topic POST-HUB-POST-TOPIC

#=============Cria mensagem topico==================
#docker exec --interactive --tty broker \
#kafka-console-producer --bootstrap-server broker:9091 \
#                       --topic CRIA_LANCAMENTO_MAIOR_TRINTA_DIAS

#=============Ler Mensagem topico=====================
#docker exec --interactive --tty broker \
#kafka-console-consumer --bootstrap-server broker:9091 \
#                       --topic CRIA_LANCAMENTO_MAIOR_TRINTA_DIAS \
#                       --from-beginning