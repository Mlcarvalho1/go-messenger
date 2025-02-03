Essa feacture tem o objetivo de pesquisar na lista de contatos por um contato especifico. 

Cenário 1: Pesquisa de usuário na lista de contatos
Scenario: Pesquisa por nome na lista de contatos
Given: O usuário “Luan Romancini” está na página “lista de contatos”
And: ve a aba de pesquisa “pesquisar por contatos”
When: pesquisa pelo nome "João" na barra de pesquisa
Then: vê todos os usuários cujo nome contém a string "João"
