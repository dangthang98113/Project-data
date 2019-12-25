package Tuan1;

import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

import javax.swing.JButton;
import javax.swing.JFrame;
import javax.swing.JTextArea;
import javax.swing.JTextField;

public class Bai2 extends JFrame implements ActionListener{
		JTextField txtNhap;
		JButton btnGen;
		JTextArea txtHien;
		public Bai2(){
			super("Primes");
			JFrame p = new JFrame();
			p.add(txtNhap = new JTextField(20));
			p.add(btnGen = new JButton("GEN"));
			p.add(txtHien= new JTextArea(30,30));
			btnGen.addActionListener(this);
			add(p);
			setSize(400,400);
			setLocation(400,200);
			setVisible(true);
		}
			public static void main(String[] args){
				new Bai2();
			}
			boolean KTNT(int n){
				int i=2;
				while(n%i != 0)
					i++;
				if(n==i)
					return true;
				return false;
				
			}
		@Override
		public void actionPerformed(ActionEvent arg0) {
			// TODO Auto-generated method stub
			int so, i, d=0;
			so = Integer.parseInt(txtNhap.getText());
			txtHien.setText("");
			boolean t=true;
			i=2;
			while(t){
				if(KTNT(i) == true){
					txtHien.append(i + "\n");
					d++;
				}
				i++;
				if(d==0)
					t=false;
			}
		}
}
