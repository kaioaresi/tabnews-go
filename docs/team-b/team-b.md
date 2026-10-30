# DRE 5 – AI Intelligent Detectors

## What

Proactively detect, investigate, and triage reliability issues by correlating Kubernetes and application observability signals with internal knowledge, creating actionable Jira tickets when intervention is required.

## Description

DRE 5 introduces an intelligent detection and automated triage capability for the DRE platform.

The capability will proactively analyze two main sources of signals:

- **Kubernetes:** cluster, namespace, and workload health, including configuration drift, probes, policies, resource conditions, and other relevant platform signals.
- **Application observability:** logs, metrics, and other available telemetry from supported services.

When DRE 5 identifies a relevant anomaly or error, it will correlate these signals and use existing internal capabilities, including **Brain AI, MCP Hub, observability tools, documentation, and approved playbooks**, to investigate the issue and determine whether action is required.

When necessary, DRE 5 will create or update a Jira ticket with the investigation context, supporting evidence, and recommended next steps for L1.

The first phase will be **read-only**, initially covering **Trino and Superset workloads running on AKS in the US region**. Automated remediation is outside the initial scope.

## Impact

DRE 5 moves part of the reliability investigation from a reactive and manual process to a proactive and evidence-driven workflow.

Instead of waiting for an alert or a reported problem, the capability can identify relevant signals, correlate infrastructure and application telemetry, and perform the initial investigation before L1 intervention.

The expected impact is:

- earlier identification of reliability issues and anomalies;
- less manual investigation and evidence collection for L1 and DRE engineers;
- more consistent triage using approved playbooks and internal knowledge;
- improved traceability through documented findings and Jira tickets;
- identification of recurring issues and observability gaps that can drive permanent reliability improvements.

Reduced downtime is an expected downstream benefit, but not a direct DRE 5 target, since resolution may depend on the team responsible for the affected workload.

## Audience

**L1 Support**  
Receives actionable Jira tickets with evidence, investigation context, and recommended next steps, reducing the effort required for initial triage.

**DRE and On-call Engineers**  
Use the investigation results to accelerate troubleshooting, identify recurring patterns, and prioritize permanent reliability improvements.

**Application and Data Workload Owners**  
Receive validated findings related to their services with enough context to investigate or correct the underlying issue.

**DRE and Data Leadership**  
Use detection, ticket quality, triage, and coverage metrics to evaluate the effectiveness of the capability and prioritize future expansion.

## Value

The primary value of DRE 5 is to reduce the time and manual effort required to move from an initial technical signal to an actionable reliability finding.

The pilot will measure:

**Detection Time**  
Time between the first detectable signal and the creation of an actionable finding or ticket for issues not already identified by existing alerts.

**Ticket Quality**  
Percentage of generated tickets that provide sufficient and accurate evidence for L1 to make a triage decision, together with the false-positive rate.

**L1 Triage Time**  
Time between ticket creation and the L1 decision to resolve, route, or reject the finding.

**Traceability**  
Percentage of validated findings linked to supporting evidence, a Jira ticket, and an identified workload owner when ownership information is available.

**Detection Coverage**  
Issues and anomalies identified by DRE 5 that were not detected by the existing monitoring and alerting mechanisms.

**Historical Validation**  
Evaluate proposed checks against confirmed historical incidents to determine whether the relevant conditions could have been detected before impact.

Baselines and numerical targets that are not currently validated will be established during the initial phase rather than treated as assumptions.

## Success

DRE 5 will be considered successful when the pilot demonstrates that proactive, read-only detection and AI-assisted investigation can consistently generate reliable and actionable findings without introducing unacceptable operational noise.

The implementation will follow controlled validation stages:

**Phase 0 – Define and Baseline**  
Confirm pilot scope, owners, supported signals, check frequency, ticket limits, KPI definitions, baseline sources, historical incidents, dependencies, and security requirements.

**Phase 1 – Shadow Mode**  
Run scheduled read-only detection and investigation without creating operational tickets. Measure findings, behavior, and running costs.

**Phase 2 – Quality Validation**  
Review findings with L1 and DRE and validate ticket quality, false positives, investigation evidence, and historical detection capability. Only validated checks can progress.

**Phase 3 – Pilot Launch**  
Enable approved checks for Trino and Superset workloads on AKS and create actionable Jira tickets when intervention is required.

**Phase 4 – Results Review**  
Compare detection time, ticket quality, triage time, traceability, detection coverage, and cost against the validated baselines and agreed targets.

**Phase 5 – Controlled Expansion**  
Expand the capability to additional services and platforms one domain at a time. Each new domain must pass the same validation process before generating operational tickets.
