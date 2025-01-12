Feature: Receber mensagem
  As a usuário
  I want to receber mensagem
  So that eu possa me comunicar.

Scenario: Receber mensagem de um contato fora do chat
    Given que estou logado na aplicação
    And meu contato "João Silva" enviou uma mensagem para mim
    When a mensagem "Tudo bem?" é recebida
    Then ela deve aparecer na minha lista de mensagens com o status "Não lida"

Scenario: Receber mensagem de um contato no chat
    Given que estou logado na aplicação
    And meu contato "João Silva" enviou uma mensagem para mim
    And estou na conversa com "João Silva"
    When a mensagem "Tudo bem?" é recebida
    Then ela deve aparecer na conversa com "João Silva"
    And "Joao Silva" deve receber a informação que a mensagem foi lida

Scenario: Receber mensagem de um grupo fora do chat
    Given que estou logado na aplicação
    And sou membro do grupo "Família"
    When o grupo "Família" recebe a mensagem "Festa no domingo às 18h"
    Then a mensagem "Festa no domingo às 18h" deve aparecer na minha lista de mensagens com o status "Não lida"

Scenario: Receber mensagem de um grupo no chat
    Given que estou logado na aplicação
    And sou membro do grupo "Família"
    And estou no chat do grupo "Família"
    When o grupo "Família" recebe a mensagem "Festa no domingo às 18h"
    Then a mensagem "Festa no domingo às 18h" deve aparecer no chat do grupo "Família"
    And todos os membros do grupo "Família" devem receber a informação que a mensagem foi lida