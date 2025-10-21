{{/*
Expand the name of the chart.
*/}}
{{- define "mk-gen.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "mk-gen.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "mk-gen.labels" -}}
helm.sh/chart: {{ include "mk-gen.chart" . }}
{{ include "mk-gen.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "mk-gen.selectorLabels" -}}
app.kubernetes.io/name: {{ include "mk-gen.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}
