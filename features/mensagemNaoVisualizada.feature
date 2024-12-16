Feature: Indicação de mensagens não visualizadas            
  As a: Usuário
  I want to: Visualizar em quais chats possuem mensagens de outros usuários não lidas por mim
  So that: Eu possa identificar mais facilmente quais chats existem mensagens de outros usuários que eu ainda não li

Scenario: O usuário possui mensagens não lidas
  Given o usuário “João” está logado
  And há mensagens não lidas no “chat com Ian”
  When o usuário “João” abrir a tela “conversas”  
  Then uma indicação visual aparece indicando que o “chat com Ian” possui mensagens não lidas
