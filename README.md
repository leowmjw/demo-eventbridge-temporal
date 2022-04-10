# demo-eventbridge-temporal
Demo of Temporal getting/sending signals via Eventbridge for a Loan 
lifecycle lasting years

cmd 

- cli - Command tool to send / receive signals
- dev - Combo independent services for easy testing

internal - Implementation for each service in whole Loan App

- user - foundation service; with sub-services: membership, profile, linked-accounts, personalization
- fraud - support service; with sub-services: identity, locale
- loan - main service; with sub-services underwriting, decision, regulations, calendar

http - Huma-powered REST/GraphQL; acts as API Gateway for each service

terraform - IaC for Eventbridge resources

webapp - Remix app

workflow - Temporal flow for each services

3rd Party Integration:

- Payment - Adyen - https://github.com/zhutik/adyen-api-go-example ; alternative: killbill.io
- Privacy - Skyflow; alternative: openprivacy.io

Business View:

- RudderStack to profile - 
- RudderStack profile from anon - https://www.rudderstack.com/blog/making-data-engineering-easier-operational-analytics-with-event-streaming-and-reverse-etl/
- PII to Redshift Serverless - https://github.com/rudderlabs/sample-user-transformers/ + https://aws.amazon.com/blogs/aws/introducing-amazon-redshift-serverless-run-analytics-at-any-scale-without-having-to-manage-infrastructure/
- Customer Journey Analysis - https://hub.getdbt.com/rudderlabs/rudder_customer_journey_analysis/latest/
- Data schema - dbt; integrate with Looker-alt: https://github.com/lightdash/lightdash
- CDP - https://rittmananalytics.com/data-centralization
- BI - LightDash - https://www.lightdash.com/lightdash-university
- snowplow-mini - w/ iqlu as event registry (https://github.com/snowplow/iglu)
- User Behavior - Indicative - https://www.indicative.com/pricing/ ; with snowplow relay
- Identity Graph - https://www.snowcatcloud.com/identity-graph/; alternatives: https://github.com/aws-samples/aws-admartech-samples/blob/master/identity-resolution/README.md