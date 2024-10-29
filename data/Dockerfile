FROM migrate/migrate

# Copy all db files
COPY ./migrations /migrations

ENTRYPOINT [ "migrate", "-path", "/migrations", "-database"]
CMD ["postgresql://postgres:123456@friend-management-postgres.db-1:5432/postgres?sslmode=disable up"]
