FLAG = "LKSPML{uncrackable_doenst_mean_unreadable}"

length = len(FLAG)

for i in range(0, length, 2):
    open(f"chall/flag{(i+1)//2}.txt", "w").write(FLAG[i] + FLAG[i + 1])