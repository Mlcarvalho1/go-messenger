Schenario : Pesquisar uma mensagem existente com sucesso
    Given eu estou na tela de conversa de um bate-papo existente
    And há mensagens no histórico da conversa
    When eu digito uma palavra-chave no campo de pesquisa
    And clico no botão de busca
    Then o sistema destaca as mensagens no histórico que contêm a palavra-chave

Schenario: Pesquisar uma mensagem inexistente
    Given eu estou na tela de conversa de um bate-papo existente
    And há mensagens no histórico da conversa
    When eu digito uma palavra-chave no campo de pesquisa
    And clico no botão de busca
    Then o sistema exibe um alerta: "Nenhuma mensagem encontrada."
    And o histórico de mensagens permanece sem alterações.
