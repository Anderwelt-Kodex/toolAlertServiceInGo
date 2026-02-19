# toolAlertServiceInGo
tool runtime check for safe usage in automated production

module package: toolservicealert
language: go lang
version: go 1.25.0

used: VisualStudioCode, Windows 11

idea: 
In automated machine procedures, tool malfunctions are a significant cost factor. The duty cycle and duration of use vary greatly across different tools, as do the manufacturer warranties—even within a single set. Since manual inspection before every production cycle is often unfeasible, an automated tracking system is essential to monitor usage and prevent critical overuse.

function:
When preparing a production session -> Select tools for session. -> Each tool gets vetted; if it will reach different thresholds. Where it might need service or be replaced.
