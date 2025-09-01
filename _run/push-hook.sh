#!/bin/bash
echo "[HOOK]" "Push"

run_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
values_dir="$run_dir/values"
script_dir="$run_dir/scripts"
root_path=$(cd "$run_dir/.." && pwd)

#############################################################################

bash "$script_dir/go_tidy_all.sh"
bash "$script_dir/go_creator_const.sh"

#############################################################################
exit 0

