
cp -r pre-commit $(pwd)/../.git/hooks/pre-commit
chmod +x $(pwd)/../.git/hooks/pre-commit

cp -r pre-push $(pwd)/../.git/hooks/pre-push
chmod +x $(pwd)/../.git/hooks/pre-push