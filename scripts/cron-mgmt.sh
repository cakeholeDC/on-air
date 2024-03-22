#!/bin/sh

usage () {
    cat <<USAGE_END
Usage:
    '$0 add "0 15 * * * echo 'hello world!'"' => add cron job
    '$0 list' => list all crontab jobs
    '$0 remove N' => remove crontab line N
USAGE_END
}

if [ -z "$1" ]; then
    usage >&2
    exit 0
fi

case "$1" in
    add)
        if [ -z "$2" ]; then
            echo "❌ You must provide a cron job-spec to 'add'"
            echo ""
            usage >&2
            exit 0
        fi

        tmpfile=$(mktemp)

        crontab -l >"$tmpfile"
        printf '%s\n' "$2" >>"$tmpfile"
        crontab "$tmpfile" && rm -f "$tmpfile"
        ;;
    list)
        crontab -l | cat -n
        ;;
    remove)
        if [ -z "$2" ]; then
            echo "❌ You must provide a crontab line number to 'remove'"
            echo ""
            usage >&2
            exit 1
        fi

        tmpfile=$(mktemp)

        crontab -l | sed -e "$2d" >"$tmpfile"
        crontab "$tmpfile" && rm -f "$tmpfile"
        ;;
    *)
        usage >&2
        exit 1
esac