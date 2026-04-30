set dotenv-load

# Compile TailwindCSS
[group('build')]
css:
    tailwindcss -i assets/styles/input.css -o static/styles/tailwind.css --minify

# Watch for changes and compile TailwindCSS
[group('dev')]
css-watch:
    tailwindcss -i assets/styles/input.css -o static/styles/tailwind.css --watch

# Run the development server
[group('dev')]
[default]
dev:
    air

# Compile the go binary
[group('build')]
build: css
    go build -o build/main .

# Run the go application
[group('dev')]
run: css
    go run .
