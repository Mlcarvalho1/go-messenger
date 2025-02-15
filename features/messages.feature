Feature: Envio de mensagem privada
    As a usuário
    I want to enviar mensagens para algum contato
    So that eu possa me comunicar.

Scenario: Enviar mensagem privada com sucesso
    Given que estou logado na aplicação
    And Entro na conversa com "Felipe torres"
    When eu envio a mensagem "Olá, como vai?" para "Felipe torres"
    Then a mensagem "Olá, como vai?" deve ser entregue para "Felipe torres"

Scenario: Enviar mensagem privada com falha na rede
    Given que estou logado na aplicação
    And Entro na conversa com "Felipe torres"
    When eu envio a mensagem "Olá, como vai?" para "Felipe torres"
    And a rede cai
    Then a menssagem "Olá, como vai?" mostrará um icone de erro
    And a menssagem "Olá, como vai?" terá um botão de reenvio
    And a mensagem "Olá, como vai?" não deve ser entregue para "Felipe torres"

Scenario: Clicar no botão de reenvio
    Given que estou logado na aplicação
    And Entro na conversa com "Felipe torres"
    When eu envio a mensagem "Olá, como vai?" para "Felipe torres"
    And a rede cai
    And eu clico no botão de reenvio
    Then a mensagem "Olá, como vai?" deve ser entregue para "Felipe torres"

Schenario : Responder a uma mensagem com sucesso
    Given eu estou na tela de conversa de um bate-papo existente
    And há mensagens exibidas no histórico da conversa
    When eu seleciono uma mensagem no histórico
    And clico no ícone de "responder"
    And digito uma nova mensagem no campo de texto
    And clico no botão de envio
    Then a nova mensagem é exibida no histórico, vinculada à mensagem respondida
    And a mensagem de resposta é marcada como enviada com sucesso.

Schenario: Responder uma mensagem e cancelar antes do envio
    Given eu estou na tela de conversa de um bate-papo existente
    And há mensagens exibidas no histórico da conversa
    When eu seleciono uma mensagem no histórico
    And clico no ícone de "responder"
    And clico no botão de cancelar antes de digitar ou enviar
    Then o sistema remove o modo de resposta
    And o campo de texto volta ao estado inicial.