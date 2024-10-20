# Cloudflare Status

Just my personal project to monitor cloudflare. Lol

# Info
This exporter is based on [Cloudflare API](https://www.cloudflarestatus.com/api). Below all the scrape component based on Cloudflare API.

- Summary
- Incidents
- Componets


# Build
Just run below command to build this exporter
```
make build
```

# Metrics
```
# All incidents status
cloudflare_exporter_incident{name="Billing Email Issues"} 0
cloudflare_exporter_incident{name="Bot Management Issues"} 0
cloudflare_exporter_incident{name="Browser Isolation issues in LAX"} 0
cloudflare_exporter_incident{name="CNI health check failures in Singapore"} 0
cloudflare_exporter_incident{name="Cloudflare Access API delays"} 0
cloudflare_exporter_incident{name="Cloudflare Billing Issues"} 0
cloudflare_exporter_incident{name="Cloudflare Blog Issues"} 0
cloudflare_exporter_incident{name="Cloudflare D1 Issues"} 0
cloudflare_exporter_incident{name="Cloudflare Dashboard UI Bug"} 0
cloudflare_exporter_incident{name="Cloudflare Dashboard and API service issues showing Load Balancing rules"} 0
cloudflare_exporter_incident{name="Cloudflare Dashboard and Cloudflare API service issues"} 0
cloudflare_exporter_incident{name="Cloudflare Dashboard service issues in Japan region"} 0
cloudflare_exporter_incident{name="Cloudflare Images Errors"} 0
cloudflare_exporter_incident{name="Cloudflare Pages — Issues with GitHub"} 0

# All components Status
cloudflare_exporter_component{name="Malé, Maldives - (MLE)"} 0
cloudflare_exporter_component{name="Manama, Bahrain - (BAH)"} 0
cloudflare_exporter_component{name="Manaus, Brazil - (MAO)"} 0
cloudflare_exporter_component{name="Manchester, United Kingdom - (MAN)"} 0
cloudflare_exporter_component{name="Mandalay, Myanmar - (MDL)"} 1
cloudflare_exporter_component{name="Manila, Philippines - (MNL)"} 1
cloudflare_exporter_component{name="Maputo, Mozambique - (MPM)"} 1
cloudflare_exporter_component{name="Marketing Site"} 0

# Summary
cloudflare_exporter_summary 1
```
# Dashboard
![image](https://github.com/user-attachments/assets/7f46223e-c69d-49ed-87dc-4822e94a8e85)

