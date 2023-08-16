getFlag() {
    local name=$1
    local count=0

    for arg in "$@"; do
        if [[ $arg == "--$name="* ]]; then
            local value=${arg#*=}
            echo "$value"
            count=1
            break
        fi
    done

    if [[ $count -eq 0 ]]; then
        echo "false"
    fi
}

case $1 in
dev)
    eval $(awk '!/^#/ && NF > 0 {print "export "  $1}' .env.dev)
    m=$(getFlag "m" $@)

    if [[ $m == "true" ]]; then
        make run_migrate
    else
        make run
    fi
    ;;

docker)
    eval $(awk '!/^#/ && NF > 0 {print "export "  $1}' .env.prod)
    b=$(getFlag "b" $@)

    if [[ $b == "true" ]]; then
        docker compose up -d --build
    else
        docker compose up
    fi
    ;;

*)
    echo "Unknown command"
    ;;
esac
