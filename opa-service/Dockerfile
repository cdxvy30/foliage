FROM openpolicyagent/opa:latest

COPY data/greetings.json /app/data/
COPY policy/*.rego /app/policy/

CMD ["run", "--server", "--bundle", "/app"]
