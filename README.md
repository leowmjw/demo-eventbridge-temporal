# demo-eventbridge-temporal
Demo of Temporal getting/sending signals via Eventbridge for a Loan 
lifecycle lasting years

cmd 

- cli - Command tool to send / receive signals
- dev - Combo independent services for easy testing

internal - Implementation for each service in whole Loan App

- user - foundation service; with sub-services: membership, profile, accounts
- fraud - support service; with sub-services: identity
- loan - main service; with sub-services underwriting, decision, regulations, calendar

3rd Party Integration:

- Payment - Adyen - https://github.com/zhutik/adyen-api-go-example ; alternative: killbill.io
- Privacy - Skyflow; alternative: openprivacy.io