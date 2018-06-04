# jiraguard
Jira automatic testing

docker build ./ --build-arg app_env=production -t manuelbcd/jiraguard

docker run --rm -p 8080:8080 manuelbcd/jiraguard
