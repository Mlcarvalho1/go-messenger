Feature: Modal de confirmação de ações 
As a Usuário
I want to Confirmar ou cancelar ações sensíveis por meio de um modal de confirmação
So that Eu possa ter a segurança de evitar decisões não intencionais ao realizar uma ação crítica

Scenario: Confirmar ação no modal de confirmação
Given o usuário está no modal “Confirmação”
When o usuário aciona o botão “Sim” 
Then a ação é confirmada 
And o usuário retorna à página anterior

Scenario: Cancelar ação no modal de confirmação
Given o usuário está no modal “Confirmação”
When o usuário aciona o botão “Não” 
Then a ação é cancelada
And o usuário retorna à página anterior

