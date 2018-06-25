
function downloadSemrel() {
    curl -L -o release https://gitlab.com/juhani/go-semrel-gitlab/uploads/222a87259f6162c1a59c8586226f61cf/release
    chmod +x release    
    ./release -v
    export SEMVER=$(./release next-version)
    echo $SEMVER 
}

function downloadGo() {
    go get gopkg.in/yaml.v2        
}

function testGo() {
    cd ./src
    go test -v
    cd ..
}

function buildGo() {
    cd ./src
    go build -v
    mv src ../ezrep
    chmod +x ../ezrep
    cd ..
}

function packageBeta() {
    tar -cvzf ezrep-$SEMVER-pre.tar.gz ezrep
}

function packageRelease() {
    tar -cvzf ezrep-$SEMVER.tar.gz ezrep            
}

function tagRelease() {
    ./release changelog
    ./release commit-and-tag CHANGELOG.md
}

function publishRelease() {
    ./release add-download -f ezrep-$SEMVER.tar.gz -d "Linux executable"
}

for arg in "$@"
do
    case "$arg" in 
        "download-go") downloadGo ;;
        "download-semrel") downloadSemrel ;;
        "test-go") testGo ;;
        "build-go") buildGo ;;
        "package-beta") packageBeta ;;
        "package-release") packageRelease ;;
        "tag-release") tagRelease ;;
        "publish-release") publishRelease ;;
        *) echo "Command: '$arg' not understood!"; exit 1
    esac

    if [ $? -ne 0 ] ; then
        exit 1
    fi
done
