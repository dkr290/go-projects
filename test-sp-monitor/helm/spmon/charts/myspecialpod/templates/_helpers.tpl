{{/*
Expand the name of the chart.
*/}}
{{- define "myspecialpod.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "myspecialpod.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "myspecialpod.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "myspecialpod.labels" -}}
helm.sh/chart: {{ include "myspecialpod.chart" . }}
{{ include "myspecialpod.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "myspecialpod.selectorLabels" -}}
app.kubernetes.io/name: {{ include "myspecialpod.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "myspecialpod.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "myspecialpod.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}


{{- define "objects" -}}
array:
  - |
  {{- range .Values.ingressclass.objects }}
    objectName: {{ .name  }}
    objectType: {{ .type }}
  {{- end }} 

{{- end -}}

{{- define "objectsazuresecrets" -}}
array:
  {{- range .Values.azuresecrets.objects }}
  - |
    objectName: {{ .name  }}
    objectType: {{ .type }}
  {{- end }} 

{{- end -}}

{{- define "objectsoauth" -}}
array:
  {{- range .Values.azuresecretsoauth.objects }}
  - |
    objectName: {{ .name  }}
    objectType: {{ .type }}
  {{- end }} 

{{- end -}}

{{- define "basicauth_objects" -}}
array:
  {{- range .Values.basicAuth.objects }}
  - |
    objectName: {{ .name  }}
    objectType: {{ .type }}
  {{- end }} 

{{- end -}}