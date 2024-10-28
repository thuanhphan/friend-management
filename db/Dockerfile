FROM migrate/migrate

# Copy all db files
COPY ./migrations /migrations

ENTRYPOINT [ "migrate", "-path", "/migrations", "-database"]
CMD ["postgresql://postgres:Password123@friend-management-go-postgres.db-1:5432/demo?sslmode=disable up"]