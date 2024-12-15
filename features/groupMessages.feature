Feature: Envio de mensagem em grupo
    As a usuário
    I want to mandar mensagens em algum grupo
    So that todos do grupo vejam minha mensagem.

Scenario: Enviar mensagem em grupo com sucesso
    Given que estou logado na aplicação
    And sou membro do grupo "Amigos de Trabalho"
    When eu envio a mensagem "Reunião às 15h" no grupo "Amigos de Trabalho"
    Then todos os membros do grupo "Amigos de Trabalho" devem receber a mensagem "Reunião às 15h"

Scenario: Enviar mensagem em grupo com falha na rede
    Given que estou logado na aplicação
    And sou membro do grupo "Amigos de Trabalho"
    When eu envio a mensagem "Reunião às 15h" no grupo "Amigos de Trabalho"
    And a rede cai
    Then a mensagem "Reunião às 15h" mostrará um icone de erro
    And a mensagem "Reunião às 15h" terá um botão de reenvio
    And a mensagem "Reunião às 15h" não deve ser entregue para "Amigos de Trabalho"
