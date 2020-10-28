{{- define "gatewaypods.labels" }}
      labels:
           app: "{{ $.Release.Name }}"
           k8s-app-name: famk8s
           owned_by: platform
{{- end }}


{{- define "gatewaydeps.labels" }}
  labels:
       app: "{{ $.Release.Name }}"
       k8s-app-name: famk8s
       owned_by: platform
{{- end }}

{{- define "deployment.apiVersion" }}apps/v1{{- end }}

{{- define "service.apiVersion" }}v1{{- end }}

{{- define "service.nameSpace" }}v1{{- end }}