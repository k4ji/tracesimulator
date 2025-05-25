# TraceSimulator

[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](./LICENSE)

---

## ⚠️ DEPRECATION NOTICE

This repository has been **archived** and is no longer actively maintained.

The tracesimulator functionality has been migrated into the
[`tracesimulationreceiver`](https://github.com/k4ji/tracesimulationreceiver) repository
under the `internal/tracesimulator` package.

If you're currently using this package (or want to use it), please open an issue in the
[`tracesimulationreceiver`](https://github.com/k4ji/tracesimulationreceiver/issues) repo
to discuss your use case.

We’ll consider re-publishing the simulator as a standalone or public package if there is demonstrated demand.

---

## Overview

**TraceSimulator** is a Go library for simulating traces by providing a high-level **blueprint** of trace structures and their relationships.

It handles:

- Generating traces with runtime details such as:
  - Trace IDs and Span IDs
  - Timestamps (based on wait and duration)
  - Parent-child relationships
  - Span links
  - Span statuses (`OK`/`ERROR`)
- Exporting spans into various formats, including:
  - **OpenTelemetry**-compatible spans

---

## Core Concepts

### Task-Based Trace Description

TraceSimulator introduces the concept of a **task** to describe spans.  
A task represents a static description of a span and includes:

- Task name
- Wait time before start and execution duration
- Relationships to other tasks (parent/child/linked)
- Metadata such as error probability, attributes, etc.

### Blueprint-Driven Trace Simulation

Traces are defined using a **blueprint**, which organizes tasks into trees.  
The blueprint is interpreted by the simulator and transformed into a runtime model of spans with realistic behavior.

You can use built-in blueprints or define your own to simulate a wide range of distributed system scenarios.  
For example, you can use `service.Blueprint` to model service interactions by grouping tasks under different services.

> See the [simulator test cases](./pkg/simulator_test.go) for examples of how to use blueprints in practice.

### Exporting to Other Formats

TraceSimulator supports exporting the simulated traces into other formats.

In particular, you can export them as **OpenTelemetry-compatible spans**, enabling:

- Visualization with tools like Jaeger, Datadog, Dynatrace, Honeycomb, etc.
- Testing and validation of tracing infrastructure
- Simulation of distributed system behavior for experiments and demos

---

## Use Cases

- Generating synthetic trace data for observability testing
- Simulating failure scenarios with controlled error probabilities
- Benchmarking backend systems with high-volume trace generation
- Building demos and training materials for tracing tools

---

## Getting Started
TODO

---

## License

Licensed under the [Apache 2.0 License](./LICENSE).
