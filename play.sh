echo "Get the current months statistics and details"
echo "$ ghkpi repo -t pr-kpi -r current  | jq ."
ghkpi repo -t pr-kpi -r current  | jq .
sleep 2
echo "Get the current months statistics with only totals"
echo "$ ghkpi repo -t pr-kpi -r current -a | jq ."
ghkpi repo -t pr-kpi -r current -a  | jq .
sleep 2
echo "Get the prior months statistics with only totals"
echo "$ ghkpi repo -t pr-kpi -r prior -a | jq ."
ghkpi repo -t pr-kpi -r prior -a | jq .
sleep 2
echo "Get statistics with a custom date range"
echo "$ ghkpi repo -t pr-kpi -s "2020-08-01T00:00:00.00Z" -e "2020-08-30T23:59:59.59Z" -a | jq ."
ghkpi repo -t pr-kpi -s "2020-08-01T00:00:00.00Z" -e "2020-08-30T23:59:59.59Z" | jq .
