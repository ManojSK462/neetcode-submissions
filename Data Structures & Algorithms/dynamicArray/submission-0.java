class DynamicArray {

    private int[] a;
    private int size;
    private int capacity;

    public DynamicArray(int capacity) {
        this.capacity = capacity;
        this.a = new int[capacity];
        this.size = 0;
    }

    public int get(int i) {
        return a[i];
    }

    public void set(int i, int n) {
        a[i] = n;
    }

    public void pushback(int n) {
        if(capacity == size){
            resize();
        }
        a[size] = n;
        size++;
    }

    public int popback() {
        size--;
        return a[size];
    }

    private void resize() {
        capacity*=2;
        int[] na = new int[capacity];
        for(int i =0;i<size;i++){
            na[i]  = a[i];
        }
        a=na;
    }

    public int getSize() {
        return size;
    }

    public int getCapacity() {
        return capacity;
    }
}
