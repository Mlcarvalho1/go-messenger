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