Feature: Lista de Contatos
As a: Usuário
I want to: Ver todos os usuários da aplicação na lista de contatos independentemente de ter conversado com eles ou não
So that: Eu possa visualizar e interagir com outros usuários cadastrados.

Scenario 1: Visualizar perfil ao clicar em um usuário
Given:  O usuário “Luan Romancini” está na página “lista de contatos”
And: ve o nome "Guilherme" na lista
When: clica no nome "Guilherme"
Then: vê as informações do perfil de "Guilherme"
